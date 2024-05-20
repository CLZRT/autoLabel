package autolabel

import (
	"context"
	"fmt"
	"log"
	"strings"

	compute "cloud.google.com/go/compute/apiv1"
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
)

// AuditLogEntry represents a LogEntry as described at
// https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry
type AuditLogEntry struct {
	ProtoPayload *AuditLogProtoPayload `json:"protoPayload"`
}

// AuditLogProtoPayload represents AuditLog within the LogEntry.protoPayload
// See https://cloud.google.com/logging/docs/reference/audit/auditlog/rest/Shared.Types/AuditLog
type AuditLogProtoPayload struct {
	MethodName         string                 `json:"methodName"`
	ResourceName       string                 `json:"resourceName"`
	AuthenticationInfo map[string]interface{} `json:"authenticationInfo"`
	Request            *request               `json:"request"`
}
type request struct {
	MachineType string `json:"machineType"`
}

var client *compute.InstancesClient

func init() {

	functions.CloudEvent("labelGceInstance", labelGceInstance)
}

// Cloud Function that receives GCE instance creation Audit Logs, and adds a
// `creator` label to the instance.
func labelGceInstance(ctx context.Context, ev event.Event) error {
	// Extract parameters from the Cloud Event and Cloud Audit Log data
	logAudit := &AuditLogEntry{}
	if err := ev.DataAs(logAudit); err != nil {
		err = fmt.Errorf("event.DataAs() : %w", err)
		log.Printf("Error parsing proto payload: %s", err)
		return err
	}
	log.Println(ev.String())
	/*

	 */
	payload := logAudit.ProtoPayload
	// compute.googleapis.com/projects/<PROJECT>/zones/<ZONE>/instances/<INSTANCE>
	paths := strings.Split(ev.Subject(), "/")

	// Get relevant VM instance details from the event's `subject` property
	// Subject format:
	if strings.Contains(logAudit.ProtoPayload.MethodName, "compute.instances.insert") && strings.Contains(payload.Request.MachineType, "machineType") {
		err := labelInstance_insert(logAudit, paths)
		if err != nil {
			log.Printf("Error labeling instance: %s", err)
			return err
		}
	} else if strings.Contains(payload.MethodName, "compute.instances.setMachineType") && strings.Contains(payload.Request.MachineType, "machineType") {
		log.Printf("setMachineType")
	} else if strings.Contains(payload.MethodName, "compute.regionInstances.bulkInsert") {
		log.Printf("bulkInsert")
	} else {
		fmt.Print("This log-message is excluded")
	}

	return nil
}
