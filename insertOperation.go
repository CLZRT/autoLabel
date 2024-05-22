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
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9_]+")
	creatorString := labelSanitizer.ReplaceAllString(strings.ToLower(creator), "_")

	var labels = map[string]string{
		"createdBy":    creatorString,
		"projectId":    resourceNameArray[2],
		"zone":         resourceNameArray[4],
		"instanceId":   payload.Response.InstanceId,
		"instanceName": resourceNameArray[6],
		"machineType":  payload.Request.MachineType,
	}

	// Set instance's label
	err := labelOperation(labels)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("The inserted instance %s has been  labeled successfully", paths[6])
	return nil

}
