package memory

import (
	"cloud.google.com/go/redis/apiv1/redispb"
	"clzrt.io/autolabel/struct/logstruct"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func RedisInstance(logAudit *logstruct.RedisLog) error {

	/**
	These field's value is needed in log file

	logAudit.ProtoPayload.ResourceName
	logAudit.ProtoPayload.AuthenticationInfo.PrincipalEmail
	logAudit.ProtoPayload.Request.Instance.MemorySizeGb

	*/
	log.Printf("get entry in GetMemortStore")
	log.Printf("ResourceName: %s", logAudit.ProtoPayload.ResourceName)
	var state string
	var err error
	var instance *redispb.Instance
	for state != "READY" {
		instance, err = GetMemoryStore(logAudit.ProtoPayload.ResourceName)
		if instance == nil || err != nil {
			return err
		}
		state = instance.State.String()

	}

	resourceNameArray := strings.Split(logAudit.ProtoPayload.ResourceName, "/")
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9-]+")
	creatorString := labelSanitizer.ReplaceAllString(strings.ToLower(logAudit.ProtoPayload.AuthenticationInfo.PrincipalEmail), "-")
	memorySize := strconv.Itoa(logAudit.ProtoPayload.Request.Instance.MemorySizeGb)
	// Set Instance's Label
	labels := map[string]string{
		"created-by":     creatorString,
		"instance-name":  resourceNameArray[5],
		"memory-size-gb": memorySize,
	}
	log.Printf("get entry in SetLabelMemoryStore")
	log.Printf("created-by: %v", labels["created-by"])

	err = SetLabelMemoryStore(instance, labels)
	if err != nil {
		return err
	}
	return nil
}
