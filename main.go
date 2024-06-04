package main

import (
	"clzrt.io/autolabel/compute/gce"
	"clzrt.io/autolabel/compute/gke"
	"clzrt.io/autolabel/database/bigquery"
	"clzrt.io/autolabel/database/memory"
	"clzrt.io/autolabel/database/sql"
	"clzrt.io/autolabel/storage/disk"
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

func main() {
	// Call the test function with the path to your JSON file
	err := TestLabelResource("./log/log.json")
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

//func init() {
//
//	functions.CloudEvent("labelResource", labelResource)
//}

func labelResource(ctx context.Context, ev event.Event) error {
	// Extract parameters from the Cloud Event and Cloud Audit Log data
	var msg logstruct.MessagePublishedData
	if err := ev.DataAs(&msg); err != nil {
		return fmt.Errorf("event.DataAs: %w", err)
	}

	/*
		decoded from base64
	*/
	logString := string(msg.Message.Data)
	/**
	todo: 1.Persistent Disk "gce_disk"
	todo: 2.FileStore "audited_resource"
	todo: 3.cloudStorage "gcs_bucket"
	todo: 4.CloudSQL "cloudsql_database"
	todo: 5.SSD
	todo: 6.MemoryStore
	todo: 7.Dataproc "gce_project"
	todo: 8.patchWork 不支持
	todo: 9.VPC Network
	todo: 10,GKE "k8s_cluster"
	todo: 11.Artifact Registry
	todo: 12.KMS
	todo: 13 GCE "gce_instance"
	*/
	// switch into which function

	methodName := gjson.Get(logString, "protoPayload.methodName").String()
	if strings.Contains(methodName, "compute") {
		log.Printf("Get into Compute")
		/*
			v1.compute.regionInstances.bulkInsert
			beta.compute.instances.insert
			v1.compute.instances.setMachineType
		*/
		if strings.Contains(methodName, "instances") || strings.Contains(methodName, "bulkInsert") {
			log.Printf("Get into Instances")
			gceLog := new(logstruct.GceLog)
			err := json.Unmarshal([]byte(logString), gceLog)
			if err != nil {
				return err
			}
			log.Printf("Get Into GCE")
			err = gce.InstanceGce(gceLog)
			if err != nil {
				return err
			}
		} else if strings.Contains(methodName, "disk") {
			diskLog := new(logstruct.DiskLog)
			err := json.Unmarshal([]byte(logString), diskLog)
			if err != nil {
				return err
			}
			err = disk.SingleDisk(diskLog)
			if err != nil {
				return err
			}
		}

	} else if strings.Contains(methodName, "sql") {
		//  cloudsql.instances.create
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

	} else if strings.Contains(methodName, "redis") {
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

	} else if strings.Contains(methodName, "bigquery") {
		log.Printf("resource Type:" + "bigquery")
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
	} else if strings.Contains(methodName, "table") {
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

	} else if strings.Contains(methodName, "container") {
		log.Printf("label gke")
		gkeLog := new(logstruct.Gkelog)
		err := json.Unmarshal([]byte(logString), gkeLog)
		if err != nil {
			return err
		}
		err = gke.GKE_Cluster(gkeLog)
		if err != nil {
			return err
		}
	} else {
		log.Printf("Excluded")
	}
	return nil

}
