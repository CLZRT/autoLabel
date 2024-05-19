package autolabel

import (
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
)

// Hello returns a greeting for the named person.
func labelInstance_insert(logString string) {
	projectId := gjson.Get(logString, "resource.labels.project_id")
	zone := gjson.Get(logString, "resource.labels.zone")
	instanceId := gjson.Get(logString, "resource.labels.instance_id")
	creatBy := strings.ReplaceAll(gjson.Get(logString, "protoPayload.authenticationInfo.principalEmail").String(), "-", "_")
	resourceNameArray := strings.Split(gjson.Get(logString, "protoPayload.resourceName").String(), "/")
	machineTypeArray := strings.Split(gjson.Get(logString, "protoPayload.request.machineType").String(), "/")

	machineType := machineTypeArray[5]
	instanceName := resourceNameArray[5]

	var labels = map[string]string{
		"createdBy":    creatBy,
		"projectId":    projectId.String(),
		"zone":         zone.String(),
		"instanceId":   instanceId.String(),
		"instanceName": instanceName,
		"machineType":  machineType,
	}
	err := labelOperation(instanceName, labels)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Label Insert Instance success")

}
