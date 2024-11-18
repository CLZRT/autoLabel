package gcs

import (
	"clzrt.io/autolabel/struct/logstruct"
	"log"
	"regexp"
	"strings"
)

func Bucket(logAudit *logstruct.Gcslog) error {
	resourceNameArray := strings.Split(logAudit.ProtoPayload.ResourceName, "/")
	bucketName := resourceNameArray[len(resourceNameArray)-1]
	bucket, err := Getbucket(bucketName)
	if err != nil {
		return err
	}
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9-]+")
	creatorString := labelSanitizer.ReplaceAllString(strings.ToLower(logAudit.ProtoPayload.AuthenticationInfo.PrincipalEmail), "-")

	labels := map[string]string{
		"created-by": creatorString,
	}
	err = SetBucketLabel(bucketName, labels, bucket)
	if err != nil {
		return err
	}
	log.Println("Bucket Set Label Success," + bucketName)
	return nil
}
