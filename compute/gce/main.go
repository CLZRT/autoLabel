package gce

import (
	"clzrt.io/autolabel/storage/disk"
	"clzrt.io/autolabel/struct/logstruct"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func NewGce(logAudit *logstruct.GceLog) error {

	resourceNameArray := strings.Split(logAudit.ProtoPayload.ResourceName, "/")
	resourceLocation := map[string]string{
		"project-id":  resourceNameArray[1],
		"zone":        resourceNameArray[3],
		"instance-id": resourceNameArray[5],
	}
	// Get instance
	log.Printf("resourceLocation: %v", resourceLocation)
	log.Printf("get entry in getInstance")
	instance, err := getInstance(resourceLocation)

	if err != nil {
		log.Println(err)
		return err
	}

	// extra info from logstruct
	operator := logAudit.ProtoPayload.AuthenticationInfo.PrincipalEmail
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9-]+")
	operatorString := labelSanitizer.ReplaceAllString(strings.ToLower(operator), "-")
	instanceId := logAudit.Resource.Labels.InstanceId
	machineTypeArray := strings.Split(instance.GetMachineType(), "/")
	// setInstanceLabel
	labelFingerprint := instance.GetLabelFingerprint()
	labels := map[string]string{}

	// 判断 实例是否存在标签
	existLabels := instance.GetLabels()
	if existLabels == nil {
		labels["created-by"] = operatorString
		labels["machine-type"] = machineTypeArray[len(machineTypeArray)-1]
		labels["instance-id"] = instanceId
		labels["instance-name"] = instance.GetName()
	} else {
		for k, v := range existLabels {
			labels[k] = v
		}
		labels["updated-by"] = operatorString
		labels["machine-type"] = machineTypeArray[len(machineTypeArray)-1]
	}
	log.Printf("labels: %v", labels)
	log.Printf("get entry in setInstanceLabel")
	err = setInstanceLabel(resourceLocation, labels, &labelFingerprint)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("The inserted instance %s has been  labeled successfully", instance.GetName())

	// Set Disk's Label
	disks := instance.GetDisks()
	for _, diskInfo := range disks {
		diskName := diskInfo.GetDeviceName()
		resourceLocation := map[string]string{
			"project-id": resourceNameArray[1],
			"zone":       resourceNameArray[3],
			"name":       instance.GetName(),
		}
		getDisk, err := disk.GetDisk(resourceLocation)
		if err != nil {
			return err
		}
		fingerprint := getDisk.GetLabelFingerprint()
		labelsDisk := map[string]string{
			"created-by":    operatorString,
			"size-gb":       strconv.FormatInt(diskInfo.GetDiskSizeGb(), 10),
			"instance-id":   instanceId,
			"instance-name": instance.GetName(),
		}
		err = disk.SetDiskLabel(resourceLocation, labelsDisk, &fingerprint)
		if err != nil {
			return err
		}
		log.Printf("The inserted instance's disk %s has been  labeled successfully", diskName)

	}

	return nil

}

func UpdateGce(logAudit *logstruct.GceLog) error {

	resourceNameArray := strings.Split(logAudit.ProtoPayload.ResourceName, "/")
	resourceLocation := map[string]string{
		"project-id":  resourceNameArray[1],
		"zone":        resourceNameArray[3],
		"instance-id": resourceNameArray[5],
	}
	// Get instance
	log.Printf("resourceLocation: %v", resourceLocation)
	log.Printf("get entry in getInstance")
	instance, err := getInstance(resourceLocation)

	if err != nil {
		log.Println(err)
		return err
	}

	// extra info from logstruct
	operator := logAudit.ProtoPayload.AuthenticationInfo.PrincipalEmail
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9-]+")
	operatorString := labelSanitizer.ReplaceAllString(strings.ToLower(operator), "-")
	instanceId := logAudit.Resource.Labels.InstanceId
	machineTypeArray := strings.Split(instance.GetMachineType(), "/")
	// setInstanceLabel
	labelFingerprint := instance.GetLabelFingerprint()
	labels := map[string]string{}

	// 判断 实例是否存在标签
	existLabels := instance.GetLabels()
	if existLabels == nil {
		labels["created-by"] = operatorString
		labels["machine-type"] = machineTypeArray[len(machineTypeArray)-1]
		labels["instance-id"] = instanceId
		labels["instance-name"] = instance.GetName()
	} else {
		for k, v := range existLabels {
			labels[k] = v
		}
		labels["updated-by"] = operatorString
		labels["machine-type"] = machineTypeArray[len(machineTypeArray)-1]
	}
	log.Printf("labels: %v", labels)
	log.Printf("get entry in setInstanceLabel")
	err = setInstanceLabel(resourceLocation, labels, &labelFingerprint)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("The inserted instance %s has been  labeled successfully", instance.GetName())

	// Set Disk's Label
	disks := instance.GetDisks()
	for _, diskInfo := range disks {
		diskName := diskInfo.GetDeviceName()
		resourceLocation := map[string]string{
			"project-id": resourceNameArray[1],
			"zone":       resourceNameArray[3],
			"name":       instance.GetName(),
		}
		getDisk, err := disk.GetDisk(resourceLocation)
		if err != nil {
			return err
		}
		fingerprint := getDisk.GetLabelFingerprint()
		labelsDisk := map[string]string{
			"created-by":    operatorString,
			"size-gb":       strconv.FormatInt(diskInfo.GetDiskSizeGb(), 10),
			"instance-id":   instanceId,
			"instance-name": instance.GetName(),
		}
		err = disk.SetDiskLabel(resourceLocation, labelsDisk, &fingerprint)
		if err != nil {
			return err
		}
		log.Printf("The inserted instance's disk %s has been  labeled successfully", diskName)

	}

	return nil

}
