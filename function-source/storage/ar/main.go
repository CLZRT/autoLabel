package ar

import (
	"clzrt.io/autolabel/struct/logstruct"
	"log"
	"regexp"
	"strings"
)

func Artifactregistry(logAudit *logstruct.Arlog) error {
	//projects/YOUR_PROJECT_ID/locations/YOUR_LOCATION/repositories/YOUR_REPOSITORY_ID
	repoName := logAudit.ProtoPayload.ResourceName + "/repositories/" + logAudit.ProtoPayload.Request.RepositoryId
	println(repoName)
	repo, err := GetRepository(repoName)
	if err != nil {
		log.Println(err)
		return err
	}
	creator := logAudit.ProtoPayload.AuthenticationInfo.PrincipalEmail
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9-]+")
	creatorString := labelSanitizer.ReplaceAllString(strings.ToLower(creator), "-")
	ArtifactregistryId := logAudit.ProtoPayload.Request.RepositoryId
	labels := map[string]string{
		"created-by":   creatorString,
		"cluster-name": ArtifactregistryId,
	}
	err = SetRepositoryLabel(repo, labels)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("The inserted instance %s has been  labeled successfully")

	return nil
}
