package gce

import (
	"autolabel/logstruct"
	"log"
	"regexp"
	"strings"
)

// construct Labels and Call Label util
func SingleInstance(logAudit *logstruct.AuditLogEntry) error {
	// Get Instance
	instance, err := getInstance(logAudit.Resource.Labels)
	if err != nil {
		log.Println(err)
		return err
	}

	// Extra Info from log
	payload := logAudit.ProtoPayload
	resourceNameArray := strings.Split(payload.ResourceName, "/")
	machineTypeArray := strings.Split(payload.Request.MachineType, "/")
	creator := payload.Response.User
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9_]+")
	creatorString := labelSanitizer.ReplaceAllString(strings.ToLower(creator), "_")

	// Set Label
	labels := map[string]string{
		"created-by":    creatorString,
		"instance-id":   payload.Response.InstanceId,
		"instance-name": resourceNameArray[5],
		"machine-type":  machineTypeArray[5],
	}
	labelFingerprint := instance.GetLabelFingerprint()
	err = setInstanceLabel(logAudit.Resource.Labels, labels, &labelFingerprint)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("The inserted instance %s has been  labeled successfully", instance.GetName())

	return nil

}

func MultiInstance(logAudit *logstruct.AuditLogEntry) error {

	// Get instance
	instance, err := getInstance(logAudit.Resource.Labels)

	if err != nil {
		log.Println(err)
		return err
	}

	// extra info from log
	creator := logAudit.ProtoPayload.AuthenticationInfo.PrincipalEmail
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9_]+")
	creatorString := labelSanitizer.ReplaceAllString(strings.ToLower(creator), "_")
	instanceId := logAudit.Resource.Labels.ResourceId

	// setLabel
	labelFingerprint := instance.GetLabelFingerprint()
	labels := map[string]string{
		"created-by":    creatorString,
		"machine-type":  instance.GetMachineType(),
		"instance-id":   instanceId,
		"instance-name": instance.GetName(),
	}
	err = setInstanceLabel(logAudit.Resource.Labels, labels, &labelFingerprint)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("The inserted instance %s has been  labeled successfully", instance.GetName())
	return nil

}
