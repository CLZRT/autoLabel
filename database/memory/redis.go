package memory

import (
	"autolabel/logstruct"
	"regexp"
	"strings"
)

func LabelInstance(logAudit *logstruct.AuditLogEntry) error {
	instance, err := GetMemoryStore(logAudit.ProtoPayload.ResourceName)
	if err != nil {
		return err
	}

	resourceNameArray := strings.Split(logAudit.ProtoPayload.ResourceName, "/")
	machineTypeArray := strings.Split(logAudit.ProtoPayload.Request.MachineType, "/")
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9-]+")
	creatorString := labelSanitizer.ReplaceAllString(strings.ToLower(logAudit.ProtoPayload.AuthenticationInfo.PrincipalEmail), "-")
	// Set Instance's Label
	labels := map[string]string{
		"created-by":    creatorString,
		"instance-id":   logAudit.ProtoPayload.Request.Instance_Id,
		"instance-name": resourceNameArray[5],
		"machine-type":  machineTypeArray[5],
	}
	err = SetLabelMemoryStore(instance, labels)
	if err != nil {
		return err
	}
	return nil
}
