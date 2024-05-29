package sql

import (
	"autolabel/logstruct"
	"autolabel/storage/disk"
	"log"
	"regexp"
	"strings"
)

func SingleDatabse(logAudit *logstruct.AuditLogEntry) error {
	// Get Instance
	instance, err := GetSql(logAudit.Resource.Labels)
	if err != nil {
		log.Println(err)
		return err
	}

	// Extra Info from log
	payload := logAudit.ProtoPayload
	resourceNameArray := strings.Split(payload.ResourceName, "/")
	machineTypeArray := strings.Split(payload.Request.Body.settings.tier, "/") //后期需要修改
	creator := payload.Response.User
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9-]+") //
	creatorString := labelSanitizer.ReplaceAllString(strings.ToLower(creator), "-")

	// Set Instance's Label
	labels := map[string]string{
		"created-by":    creatorString,
		"database-id":   resource.labels.database_id
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
			InstanceId: diskName,
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
