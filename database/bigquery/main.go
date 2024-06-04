package bigquery

import (
	"clzrt.io/autolabel/struct/logstruct"
	"log"
	"regexp"
	"strings"
)

func BigQueryDataset(logAudit *logstruct.DatasetlogBg) error {
	resourceLocation := map[string]string{
		"project-id": logAudit.Resource.Labels.ProjectId,
		"dataset-id": logAudit.Resource.Labels.DatasetId,
	}
	for k, v := range resourceLocation {
		log.Printf(k, "value is", v)
	}
	log.Printf("resourceLocation: %v", resourceLocation)
	log.Printf("get entry in GetDataset ")
	dataset, err := GetDataset(resourceLocation)
	if err != nil {
		return err
	}

	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9-]+")
	creatorString := labelSanitizer.ReplaceAllString(strings.ToLower(logAudit.ProtoPayload.AuthenticationInfo.PrincipalEmail), "-")
	labels := map[string]string{
		"created-by": creatorString,
		"dataset-id": logAudit.Resource.Labels.DatasetId,
	}
	log.Printf("get entry in SetDatasetLabel")
	err = SetDatasetLabel(labels, dataset)
	if err != nil {
		return err
	}
	return nil
}

func BigQueryTable(logAudit *logstruct.TablelogBG) error {
	tableLocation := logAudit.ProtoPayload.ServiceData.TableInsertRequest.Resource.TableName
	resourceLocation := map[string]string{
		"project-id": tableLocation.ProjectId,
		"table-id":   tableLocation.TableId,
		"dataset-id": tableLocation.DatasetId,
	}
	table, err := GetTable(resourceLocation)
	if err != nil {
		return err
	}
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9-]+")
	creatorString := labelSanitizer.ReplaceAllString(strings.ToLower(logAudit.ProtoPayload.AuthenticationInfo.PrincipalEmail), "-")

	label := map[string]string{
		"created-by": creatorString,
		"dataset-id": tableLocation.DatasetId,
		"table-id":   tableLocation.TableId,
	}

	err = SetTableLabel(label, table)
	if err != nil {
		return err
	}
	return nil
}
