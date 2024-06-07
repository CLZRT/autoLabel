package autolabel

import (
	"clzrt.io/autolabel/compute/dataproc"
	"clzrt.io/autolabel/compute/gce"
	"clzrt.io/autolabel/compute/gke"
	"clzrt.io/autolabel/compute/ipaddress"
	"clzrt.io/autolabel/database/bigquery"
	"clzrt.io/autolabel/database/memory"
	"clzrt.io/autolabel/database/sql"
	"clzrt.io/autolabel/devops/deploy"
	"clzrt.io/autolabel/security/apigateway"
	"clzrt.io/autolabel/storage/ar"
	"clzrt.io/autolabel/storage/disk"
	"clzrt.io/autolabel/storage/filestore"
	"clzrt.io/autolabel/storage/gcs"
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
			if strings.Contains(methodName, "instances.insert") {
				if gjson.Get(logString, "operation.last").String() == "true" {
					gceLog := new(logstruct.GceLog)
					err := json.Unmarshal([]byte(logString), gceLog)
					if err != nil {
						return err
					}
					err = gce.NewGce(gceLog)
					if err != nil {
						return err
					}
				}

			} else if strings.Contains(methodName, "instances.setMachineType") {
				gceLog := new(logstruct.GceLog)
				err := json.Unmarshal([]byte(logString), gceLog)
				if err != nil {
					return err
				}
				err = gce.UpdateGce(gceLog)
				if err != nil {
					return err
				}
			}

		} else if strings.Contains(methodName, "bulkInsert") {
			log.Printf("Get Into Instance")
			gceLog := new(logstruct.GceLog)
			err := json.Unmarshal([]byte(logString), gceLog)
			if err != nil {
				return err
			}
			log.Printf("Get Into GCE")
			err = gce.BulkGce(gceLog)
			if err != nil {
				return err
			}
		} else if strings.Contains(methodName, "disks.insert") {
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
		} else if strings.Contains(methodName, "addresses.insert") {
			addressLog := new(logstruct.IpaddressLog)
			err := json.Unmarshal([]byte(logString), addressLog)
			if err != nil {
				return err
			}
			err = ipaddress.StaticIp(addressLog)
			if err != nil {
				return err
			}
		} else if strings.Contains(methodName, "globalAddresses.insert") {
			addressLog := new(logstruct.GlobalAddressLog)
			err := json.Unmarshal([]byte(logString), addressLog)
			if err != nil {
				return err
			}
			err = ipaddress.GlobalStaticIp(addressLog)
			if err != nil {
				return err
			}
		}
	case "cloudsql.googleapis.com":
		if gjson.Get(logString, "operation.last").String() == "true" {
			if strings.Contains(methodName, "instances.create") || strings.Contains(methodName, "instances.update") {
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
		}
	case "redis.googleapis.com":
		if strings.Contains(methodName, "CreateInstance") || strings.Contains(methodName, "UpdateInstance") {
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

		if strings.Contains(methodName, "InsertDataset") {
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
		} else if strings.Contains(methodName, "insert") {
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
		if strings.Contains(methodName, "Cluster") {
			log.Printf("label dataproc cluster")
			clusterLog := new(logstruct.ClusterlogDP)
			err := json.Unmarshal([]byte(logString), clusterLog)
			if err != nil {
				return err
			}
			err = dataproc.Cluster(clusterLog)
			if err != nil {
				return err
			}
		} else if strings.Contains(methodName, "Job") {
			log.Printf("label dataproc job")
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
	case "storage.googleapis.com":
		if strings.Contains(methodName, "buckets.create") {
			log.Printf("label cloud-storage bucket")
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
	case "container.googleapis.com":
		if strings.Contains(methodName, "CreateCluster") {
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
		}
	case "artifactregistry.googleapis.com":
		if strings.Contains(methodName, "CreateRepository") {
			log.Printf("label artifactregistry")
			arLog := new(logstruct.Arlog)
			err := json.Unmarshal([]byte(logString), arLog)
			if err != nil {
				return err
			}
			err = ar.Artifactregistry(arLog)
			if err != nil {
				return err
			}
		}
	case "apigateway.googleapis.com":
		if strings.Contains(methodName, "CreateApi") {
			log.Printf("label apis")
			apislog := new(logstruct.ApigatewayLog)
			err := json.Unmarshal([]byte(logString), apislog)
			if err != nil {
				return err
			}
			err = apigateway.Api(apislog)
			if err != nil {
				return err
			}
		}
	case "clouddeploy.googleapis.com":
		if strings.Contains(methodName, "CreateTarget") {
			log.Printf("label cloud-deploy target")
			targetlog := new(logstruct.TargetLog)
			err := json.Unmarshal([]byte(logString), targetlog)
			if err != nil {
				return err
			}
			err = deploy.Target(targetlog)
			if err != nil {
				return err
			}

		} else if strings.Contains(methodName, "CreateRollout") {
			log.Printf("update cloud-deploy target label")
			rolloutlog := new(logstruct.RolloutLog)
			err := json.Unmarshal([]byte(logString), rolloutlog)
			if err != nil {
				return err
			}
			err = deploy.Rollout(rolloutlog)
			if err != nil {
				return err
			}
		}
	}

	return nil

}
