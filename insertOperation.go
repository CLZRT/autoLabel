package autolabel

import (
	"log"
	"regexp"
	"strings"
)

// construct Labels and Call Label util
func insertInstanceLabel(logAudit *AuditLogEntry) error {

	payload := logAudit.ProtoPayload
	creator := payload.Response.User
	resourceNameArray := strings.Split(payload.ResourceName, "/")
	machineTypeArray := strings.Split(payload.Request.MachineType, "/")
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9_]+")
	creatorString := labelSanitizer.ReplaceAllString(strings.ToLower(creator), "_")
	log.Printf("creator:" + creatorString)
	log.Printf("resourceNameArray:" + payload.ResourceName)
	log.Printf("instanceName:" + resourceNameArray[5])
	log.Printf("instanceId:" + payload.Response.InstanceId)
	log.Printf("machine-type:" + payload.Request.MachineType)

	labels := map[string]string{
		"created-by":    creatorString,
		"project-id":    resourceNameArray[1],
		"zone":          resourceNameArray[3],
		"instance-id":   payload.Response.InstanceId,
		"instance-name": resourceNameArray[5],
		"machine-type":  machineTypeArray[5],
	}
	log.Printf("instance-name" + labels["instance-name"])

	// Set instance's label
	err := labelOperation(labels)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("The inserted instance %s has been  labeled successfully", resourceNameArray[5])
	return nil

}
