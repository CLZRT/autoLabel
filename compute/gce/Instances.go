package gce

import (
	"autolabel/logstruct"
	"autolabel/storage/disk"
	"log"
	"regexp"
	"strings"
)

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
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9-]+")
	creatorString := labelSanitizer.ReplaceAllString(strings.ToLower(creator), "-")

	// Set Instance's Label
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

	// Set Disk's Label
	disks := instance.GetDisks()
	for _, diskInfo := range disks {
		diskName := diskInfo.GetDeviceName()

		resourceLabels := logstruct.AuditResourceLabels{
			ProjectId:  logAudit.Resource.Labels.ProjectId,
			Zone:       logAudit.Resource.Labels.Zone,
			ResourceId: diskName,
		}
		getDisk, err := disk.GetDisk(&resourceLabels)
		if err != nil {
			return err
		}
		fingerprint := getDisk.GetLabelFingerprint()
		err = disk.SetDiskLabel(&resourceLabels, labels, &fingerprint)
		if err != nil {
			return err
		}

	}
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
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9-]+")
	creatorString := labelSanitizer.ReplaceAllString(strings.ToLower(creator), "-")
	instanceId := logAudit.Resource.Labels.ResourceId

	// setInstanceLabel
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

	// Set Disk's Label
	disks := instance.GetDisks()
	for _, diskInfo := range disks {
		diskName := diskInfo.GetDeviceName()

		resourceLabels := logstruct.AuditResourceLabels{
			ProjectId:  logAudit.Resource.Labels.ProjectId,
			Zone:       logAudit.Resource.Labels.Zone,
			ResourceId: diskName,
		}
		getDisk, err := disk.GetDisk(&resourceLabels)
		if err != nil {
			return err
		}
		fingerprint := getDisk.GetLabelFingerprint()
		err = disk.SetDiskLabel(&resourceLabels, labels, &fingerprint)
		if err != nil {
			return err
		}
		log.Printf("The inserted instance's disk %s has been  labeled successfully", diskName)

	}
	return nil

}
