package log

import (
	"clzrt.io/autolabel/compute/dataproc"
	"clzrt.io/autolabel/compute/gce"
	"clzrt.io/autolabel/database/bigquery"
	"clzrt.io/autolabel/database/memory"
	"clzrt.io/autolabel/database/sql"
	"clzrt.io/autolabel/storage/disk"
	"clzrt.io/autolabel/struct/logstruct"
	"context"
	"encoding/json"
	"fmt"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/tidwall/gjson"
	"log"
	"strings"
)

func init() {

	functions.CloudEvent("labelResource", labelResource)
}

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

	} else if strings.Contains(methodName, "dataproc") {
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

	} else {
		log.Printf("Excluded")
	}

	return nil

}
