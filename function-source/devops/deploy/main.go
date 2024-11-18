package deploy

import (
	"clzrt.io/autolabel/struct/logstruct"
	"log"
	"regexp"
	"strings"
)

func Target(logAudit *logstruct.TargetLog) error {
	resourceName := logAudit.ProtoPayload.ResourceName
	log.Println("entry into get target function")
	target, err := getTarget(resourceName)
	if err != nil {
		return err
	}
	log.Println("get target: " + target.Name + " success")
	labels := map[string]string{}
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9-]+")
	operatorString := labelSanitizer.ReplaceAllString(strings.ToLower(logAudit.ProtoPayload.AuthenticationInfo.PrincipalEmail), "-")

	if target.Labels != nil {
		labels = target.Labels
		labels["updated-by"] = operatorString
	} else {
		labels["created-by"] = operatorString
		labels["target-uid"] = target.GetUid()
		labels["target-id"] = target.GetTargetId()
	}
	log.Println("entry into set target function")
	err = setTarget(labels, target)
	if err != nil {
		return err
	}
	log.Println("set target: " + target.Name + " success")
	return nil
}

func Rollout(logAudit *logstruct.RolloutLog) error {
	resourceName := logAudit.ProtoPayload.ResourceName
	log.Println("entry into get rollout2target function")
	target, err := getTargetfromRollout(resourceName)
	if err != nil {
		return err
	}
	log.Println("get target from rollout: " + target.Name + " success")
	labels := map[string]string{}
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9-]+")
	operatorString := labelSanitizer.ReplaceAllString(strings.ToLower(logAudit.ProtoPayload.AuthenticationInfo.PrincipalEmail), "-")

	if target.Labels != nil {
		//labels = target.Labels
		labels["updated-by"] = operatorString
	} else {
		labels["created-by"] = operatorString
		labels["target-uid"] = target.GetUid()
		labels["target-id"] = target.GetTargetId()
	}
	log.Println("entry into set target function")
	err = setTarget(labels, target)
	if err != nil {
		return err
	}
	log.Println("set target: " + target.Name + " success")
	return nil

}
