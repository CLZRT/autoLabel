package main

import (
	"clzrt.io/autolabel/compute/dataproc"
	"clzrt.io/autolabel/compute/gce"
	"clzrt.io/autolabel/compute/ipaddress"
	"clzrt.io/autolabel/database/bigquery"
	"clzrt.io/autolabel/database/memory"
	"clzrt.io/autolabel/database/sql"
	"clzrt.io/autolabel/storage/disk"
	"clzrt.io/autolabel/storage/filestore"
	"clzrt.io/autolabel/storage/gcs"
	"clzrt.io/autolabel/struct/logstruct"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"strings"
)

//
//func init() {
//
//	functions.CloudEvent("labelResource", labelResource)
//}

func main() {
	// Call the test function with the path to your JSON file
	err := TestLabelResource("./log/filestore_backup.json")
	if err != nil {
		log.Printf("Test failed: %v", err)
	}
}

func TestLabelResource(filePath string) error {
	// 读取 JSON 文件内容
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
		return err
	}

	// 构建 Cloud Event
	e := event.New()
	e.SetData(event.ApplicationJSON, []byte(data))

	// 构建模拟的消息结构体，适配您现有的数据结构
	msg := logstruct.MessagePublishedData{
		Message: logstruct.PubSubMessage{
			Data: []byte(data), // 这里需要的是 []byte 类型
		},
	}

	// 将 msg 对象编码为 JSON
	msgData, err := json.Marshal(msg)
	if err != nil {
		log.Fatalf("Error marshalling msg to JSON: %v", err)
		return err
	}

	// 手动设置 Event 的 Data 为 msgData
	err = e.SetData(event.ApplicationJSON, msgData)
	if err != nil {
		log.Printf("Failed to set event data: %v", err)
		return err
	}

	// 创建上下文
	ctx := context.Background()

	// 调用 labelResource 函数
	err = labelResource(ctx, e)
	if err != nil {
		log.Printf("Error during labelResource: %v", err)
		return err
	}

	return nil
}
func labelResource(ctx context.Context, ev event.Event) error {
	// Extract parameters from the Cloud Event and Cloud Audit Log data
	var msg logstruct.MessagePublishedData
	if err := ev.DataAs(&msg); err != nil {
		return fmt.Errorf("event.DataAs: %w", err)
	}

	//decoded from base64
	logString := string(msg.Message.Data)

	// switch into which function
	serviceName := gjson.Get(logString, "protoPayload.serviceName").String()
	methodName := gjson.Get(logString, "protoPayload.methodName").String()
	switch serviceName {
	case "compute.googleapis.com":
		log.Printf("Get into Compute")
		if strings.Contains(methodName, "instance") {

			log.Printf("Get Into Instance")
			gceLog := new(logstruct.GceLog)
			err := json.Unmarshal([]byte(logString), gceLog)
			if err != nil {
				return err
			}
			err = gce.InstanceGce(gceLog)
			if err != nil {
				return err
			}
		} else if strings.Contains(methodName, "bulkInsert") {
			log.Printf("Get Into Instance")
			gceLog := new(logstruct.GceLog)
			err := json.Unmarshal([]byte(logString), gceLog)
			if err != nil {
				return err
			}
			err = gce.InstanceGce(gceLog)
			if err != nil {
				return err
			}
		} else if strings.Contains(methodName, "disk") {
			log.Printf("Get into Disk")
			diskLog := new(logstruct.DiskLog)
			err := json.Unmarshal([]byte(logString), diskLog)
			if err != nil {
				return err
			}
			err = disk.SingleDisk(diskLog)
			if err != nil {
				return err
			}
		} else if strings.Contains(methodName, "addresses") {
			addressLog := new(logstruct.IpaddressLog)
			err := json.Unmarshal([]byte(logString), addressLog)
			if err != nil {
				return err
			}
			err = ipaddress.StaticIp(addressLog)
			if err != nil {
				return err
			}
		}
	case "cloudsql.googleapis.com":
		if gjson.Get(logString, "operation.last").String() == "true" {
			// instance create Complete
			sqlLog := new(logstruct.SqlLog)
			err := json.Unmarshal([]byte(logString), sqlLog)
			if err != nil {
				return err
			}
			err = sql.Database(sqlLog)
			if err != nil {
				return err
			}
		}
	case "redis.googleapis.com":
		if strings.Contains(methodName, "redis") {
			// google.cloud.redis.v1.CloudRedis.CreateInstance
			redisLog := new(logstruct.RedisLog)
			err := json.Unmarshal([]byte(logString), redisLog)
			if err != nil {
				return err
			}
			err = memory.RedisInstance(redisLog)
			if err != nil {
				return err
			}

		}
	case "bigquery.googleapis.com":
		if strings.Contains(methodName, "bigquery") {
			if strings.Contains(methodName, "Dataset") {
				log.Printf("Label dataset")
				datasetLog := new(logstruct.DatasetlogBg)
				err := json.Unmarshal([]byte(logString), datasetLog)
				if err != nil {
					return err
				}
				err = bigquery.BigQueryDataset(datasetLog)
				if err != nil {
					return err
				}
			}
		}
	case "table.googleapis.com":
		if strings.Contains(methodName, "table") {
			log.Printf("Label table")
			tableLog := new(logstruct.TablelogBG)
			err := json.Unmarshal([]byte(logString), tableLog)
			if err != nil {
				return err
			}
			err = bigquery.BigQueryTable(tableLog)
			if err != nil {
				return err
			}

		}
	case "dataproc.googleapis.com":
		if strings.Contains(methodName, "dataproc") {
			log.Printf("resource Type:" + "dataproc")
			if strings.Contains(methodName, "Cluster") {
				clusterLog := new(logstruct.ClusterlogDP)
				err := json.Unmarshal([]byte(logString), clusterLog)
				if err != nil {
					return err
				}
				err = dataproc.DataprocCluster(clusterLog)
				if err != nil {
					return err
				}
			} else if strings.Contains(methodName, "Job") {
				jobLog := new(logstruct.JoblogDP)
				err := json.Unmarshal([]byte(logString), jobLog)
				if err != nil {
					return err
				}
				err = dataproc.DataprocJob(jobLog)
				if err != nil {
					return err
				}
			}

		}
	case "storage.googleapis.com":
		if strings.Contains(methodName, "storage") {
			log.Printf("resource Type:" + "storage")
			if strings.Contains(methodName, "bucket") {
				log.Printf("Label bucket")
				gcsLog := new(logstruct.Gcslog)
				err := json.Unmarshal([]byte(logString), gcsLog)
				if err != nil {
					return err
				}
				err = gcs.Bucket(gcsLog)
				if err != nil {
					return err
				}
			}
		}
	case "file.googleapis.com":
		if strings.Contains(methodName, "CreateInstance") {
			log.Println("Label filestore's instance")
			instanceLog := new(logstruct.FilestoreInstanceLog)
			err := json.Unmarshal([]byte(logString), instanceLog)
			if err != nil {
				return err
			}
			err = filestore.FilestoreInstance(instanceLog)
			if err != nil {
				return err
			}
		}
		if strings.Contains(methodName, "CreateBackup") {
			log.Println("Label filestore's backup")
			backupLog := new(logstruct.FilestoreBackupLog)
			err := json.Unmarshal([]byte(logString), backupLog)
			if err != nil {
				return err
			}
			err = filestore.FilestoreBackup(backupLog)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
