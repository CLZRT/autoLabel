package main

import (
	"autolabel/compute/gce"
	"autolabel/logstruct"
	"context"
	"encoding/json"
	"fmt"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
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
		err = fmt.Errorf("event.DataAs() : %w", err)
		log.Printf("Error parsing proto payload: %s", err)
		return err
	}

	/*
		Load the logInfo in
	*/
	// Automatically decoded from base64.
	logInfo := new(logstruct.AuditLogEntry)
	err := json.Unmarshal(msg.Message.Data, logInfo)
	if err != nil {
		log.Printf("Error parsing proto payload: %s", err)
	}
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
	switch logInfo.Resource.Type {
	case "gce_instance":
		err := labelGceInstance(ev)
		if err != nil {
			return err
		}

		//
		//

	}
	return nil

}

// Cloud Function that receives GCE instance creation Audit Logs, and adds a
// `creator` label to the instance.
func labelGceInstance(ev event.Event) error {

	// Extract parameters from the Cloud Event and Cloud Audit Log data
	var msg logstruct.MessagePublishedData
	if err := ev.DataAs(&msg); err != nil {
		err = fmt.Errorf("event.DataAs() : %w", err)
		log.Printf("Error parsing proto payload: %s", err)
		return err
	}

	/*
		Load the logInfo in structure
	*/

	log.Printf("Log entry data: %s", string(msg.Message.Data)) // Automatically decoded from base64.
	logInfo := new(logstruct.AuditLogEntry)
	err := json.Unmarshal(msg.Message.Data, logInfo)
	if err != nil {
		log.Printf("Error parsing proto payload: %s", err)
	}

	// switch into which function
	methodArray := strings.Split(logInfo.ProtoPayload.MethodName, ".")
	switch methodArray[len(methodArray)-1] {
	case "insert":
		if logInfo.ProtoPayload.Response == nil {
			log.Printf("Excluded this message, cause no response")
			return nil
		} else {
			err := gce.SingleInstance(logInfo)
			if err != nil {
				log.Printf("insert: Error labeling instance: %s", err)
				return err
			}
		}
	case "setMachineType":
		if logInfo.ProtoPayload.Response == nil {
			log.Printf("Excluded this message, cause no response")
			return nil
		} else {
			err := gce.SingleInstance(logInfo)
			if err != nil {
				log.Printf("setMachineType: Error labeling instance: %s", err)
				return err
			}
		}
	case "bulkInsert":
		if logInfo.Resource.Labels.ResourceId == "" {
			log.Printf("Excluded this message, cause no instanceId")
			return nil
		} else {
			err := gce.MultiInstance(logInfo)
			if err != nil {
				log.Printf("bulkInsert: Error labeling instance: %s", err)
				return err
			}
		}

	}

	return nil
}
