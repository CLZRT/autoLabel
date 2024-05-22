package autolabel

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
	"log"
)

// MessagePublishedData contains the full Pub/Sub message
// See the documentation for more details:
// https://cloud.google.com/eventarc/docs/cloudevents#pubsub
type MessagePublishedData struct {
	Message PubSubMessage
}

// PubSubMessage is the payload of a Pub/Sub event.
// See the documentation for more details:
// https://cloud.google.com/pubsub/docs/reference/rest/v1/PubsubMessage
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// AuditLogEntry represents a LogEntry as described at
// https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry
type AuditLogEntry struct {
	ProtoPayload *AuditLogProtoPayload `json:"protoPayload"`
}

// AuditLogProtoPayload represents AuditLog within the LogEntry.protoPayload
// See https://cloud.google.com/logging/docs/reference/audit/auditlog/rest/Shared.Types/AuditLog
type AuditLogProtoPayload struct {
	MethodName   string         `json:"methodName"`
	Request      *AuditRequest  `json:"request"`
	ResourceName string         `json:"resourceName"`
	Response     *AuditResponse `json:"response"`
}
type AuditRequest struct {
	MachineType string `json:"machineType"`
	Name        string `json:"name"`
}
type AuditResponse struct {
	InstanceId    string `json:"targetId"`
	User          string `json:"user"`
	OperationType string `json:"operationType"`
}

func init() {

	functions.CloudEvent("labelGceInstance", labelGceInstance)
}

// Cloud Function that receives GCE instance creation Audit Logs, and adds a
// `creator` label to the instance.
func labelGceInstance(ctx context.Context, ev event.Event) error {

	// Extract parameters from the Cloud Event and Cloud Audit Log data
	var msg MessagePublishedData
	if err := ev.DataAs(&msg); err != nil {
		err = fmt.Errorf("event.DataAs() : %w", err)
		log.Printf("Error parsing proto payload: %s", err)
		return err
	}

	/*
		Load the logInfo in structure
	*/

	log.Printf("Log entry data: %s", string(msg.Message.Data)) // Automatically decoded from base64.
	logInfo := new(AuditLogEntry)
	err := json.Unmarshal(msg.Message.Data, logInfo)
	if err != nil {
		log.Printf("Error parsing proto payload: %s", err)
	}

	// switch into which function
	if logInfo.ProtoPayload.Response == nil {
		log.Printf("Excluded this message, cause no response")
		return nil
	} else {
		switch logInfo.ProtoPayload.Response.OperationType {
		case "insert":
			err := insertInstanceLabel(logInfo)
			if err != nil {
				log.Printf("Error labeling instance: %s", err)
				return err
			}
		case "setMachineType":
			err := labelInstance_update(logInfo)
			if err != nil {
				log.Printf("Error labeling instance: %s", err)
				return err
			}
		case "bulkInsert":
			err := labelInstance_bulkInsert(logInfo)
			if err != nil {
				log.Printf("Error labeling instance: %s", err)
				return err
			}

		}
	}

	return nil
}
