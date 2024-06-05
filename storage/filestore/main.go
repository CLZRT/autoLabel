package filestore

import (
	"clzrt.io/autolabel/struct/logstruct"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func FilestoreInstance(logAudit *logstruct.FilestoreInstanceLog) error {
	resourceName := logAudit.ProtoPayload.ResourceName
	log.Println("get in to GetInstance")
	instance, err := GetInstance(resourceName)
	if err != nil {
		return err
	}
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9-]+")
	creatorString := labelSanitizer.ReplaceAllString(strings.ToLower(logAudit.ProtoPayload.AuthenticationInfo.PrincipalEmail), "-")
	instanceNameArray := strings.Split(instance.GetName(), "/")
	labels := map[string]string{
		"created-by":    creatorString,
		"instance-name": instanceNameArray[len(instanceNameArray)-1],
		"instance-tier": strings.ToLower(instance.GetTier().String()),
	}
	log.Println("get in to SetInstanceLabel")
	err = SetInstanceLabel(labels, instance)
	if err != nil {
		return err
	}
	log.Println("set in to filestore instance label successfully")
	return nil
}

func FilestoreBackup(logAudit *logstruct.FilestoreBackupLog) error {
	resourceName := logAudit.ProtoPayload.ResourceName
	log.Println("get in to GetBackup")
	backup, err := GetBackup(resourceName)
	if err != nil {
		return err
	}
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9-]+")
	creatorString := labelSanitizer.ReplaceAllString(strings.ToLower(logAudit.ProtoPayload.AuthenticationInfo.PrincipalEmail), "-")
	backupNameArray := strings.Split(backup.GetName(), "/")
	sourceInstanceNameArray := strings.Split(backup.SourceInstance, "/")
	labels := map[string]string{
		"created-by":      creatorString,
		"backup-name":     backupNameArray[len(backupNameArray)-1],
		"capacity-gb":     strconv.Itoa(int(backup.CapacityGb)),
		"source-instance": sourceInstanceNameArray[len(sourceInstanceNameArray)-1],
	}

	log.Println("get in to SetBackupLabel")
	err = SetBackupLabel(labels, backup)
	if err != nil {
		return err
	}
	log.Println("set in to filestore backup label successfully")
	return nil
}
