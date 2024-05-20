package autolabel

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

// Hello returns a greeting for the named person.
func labelInstance_insert(logAudit *AuditLogEntry, paths []string) error {

	payload := logAudit.ProtoPayload
	creator, ok := payload.AuthenticationInfo["principalEmail"]
	if !ok {
		err := fmt.Errorf("principalEmail not found in payload: %v", payload)
		log.Printf("creator email not found: %s", err)
		return err
	}
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9_]+")
	creatorString := labelSanitizer.ReplaceAllString(strings.ToLower(creator.(string)), "_")

	var labels = map[string]string{
		"createdBy":    creatorString,
		"projectId":    paths[2],
		"zone":         paths[4],
		"instanceId":   paths[6],
		"instanceName": paths[6],
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
