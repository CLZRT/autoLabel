package disk

import (
	"clzrt.io/autolabel/struct/logstruct"
	"log"
	"regexp"
	"strings"
)

func SingleDisk(logAudit *logstruct.DiskLog) error {
	resourceLocation := map[string]string{
		"project-id": logAudit.Resource.Labels.ProjectId,
		"zone":       logAudit.Resource.Labels.Zone,
		"name":       logAudit.ProtoPayload.Request.Name,
	}

	disk, err := GetDisk(resourceLocation)
	if err != nil {
		log.Println(err)
		return err
	}

	//extra info from log
	creator := logAudit.ProtoPayload.AuthenticationInfo.PrincipalEmail
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9-]+")
	creatorString := labelSanitizer.ReplaceAllString(strings.ToLower(creator), "-")
	size := logAudit.ProtoPayload.Request.SizeGb
	diskId := logAudit.ProtoPayload.Response.Id
	diskName := logAudit.ProtoPayload.Request.Name
	diskTypeArray := strings.Split(*disk.Type, "/")
	diskType := diskTypeArray[len(diskTypeArray)-1]

	fingerprint := disk.GetLabelFingerprint()
	labels := map[string]string{
		"created-by": creatorString,
		"disk-size":  size,
		"disk-id":    diskId,
		"disk-name":  diskName,
		"disk-type":  diskType,
	}
	err = SetDiskLabel(resourceLocation, labels, &fingerprint)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
