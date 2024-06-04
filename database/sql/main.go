package sql

import (
	"clzrt.io/autolabel/struct/logstruct"
	"log"
	"regexp"
	"strings"
)

func Database(logAudit *logstruct.SqlLog) error {

	resourceAttributesArray := strings.Split(logAudit.ProtoPayload.AuthorizationInfo[0].ResourceAttributes.Name, "/")
	resourceLocation := map[string]string{
		"project-id":    logAudit.Resource.Labels.ProjectId,
		"database-name": resourceAttributesArray[3],
	}
	for k, v := range resourceLocation {
		log.Printf(k, "value is", v)
	}

	// Get database Instance
	instance, err := GetDatabase(resourceLocation)
	if err != nil {
		return err
	}
	// Construct label struct
	creator := logAudit.ProtoPayload.AuthenticationInfo.PrincipalEmail
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9-]+")
	creatorString := labelSanitizer.ReplaceAllString(strings.ToLower(creator), "-")
	machineType := instance.Settings.Tier
	databaseName := resourceAttributesArray[3]
	// setInstanceLabel

	labels := map[string]string{
		"created-by":    creatorString,
		"machine-type":  machineType,
		"database-name": databaseName,
	}
	for k, v := range labels {
		log.Printf(k+"'s value is", v)
	}
	err = SetDatabaseLabel(resourceLocation, labels)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}
