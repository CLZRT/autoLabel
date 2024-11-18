package gke

import (
	"clzrt.io/autolabel/struct/logstruct"
	"log"
	"regexp"
	"strings"
)

func GKE_Cluster(logAudit *logstruct.Gkelog) error {
	resourceNameArray := strings.Split(logAudit.ProtoPayload.ResourceName, "/")
	//   "resourceName": "projects/testftl-5-7/zones/us-central1-c/clusters/cluster-1-czm",
	resourceLocation := map[string]string{
		"project-id": resourceNameArray[1],
		"zone":       resourceNameArray[3],
		"cluster-id": resourceNameArray[5],
	}
	//     Specified in the format `projects/*/locations/*/clusters/*`.
	name := logAudit.ProtoPayload.Request.Parent + "/clusters/" + logAudit.ProtoPayload.Request.Cluster.Name
	log.Printf("NAME log: %v", name)
	log.Printf("get entry in gke")
	gke, err := GetGke(name, resourceLocation)

	if err != nil {
		log.Println(err)
		return err
	}

	// extra info from logstruct
	creator := logAudit.ProtoPayload.AuthenticationInfo.PrincipalEmail
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9-]+")
	creatorString := labelSanitizer.ReplaceAllString(strings.ToLower(creator), "-")
	clusterName := logAudit.Resource.Labels.ClusterName

	// setInstanceLabel
	labelFingerprint := gke.GetLabelFingerprint()

	labels := map[string]string{
		"created-by":   creatorString,
		"cluster-name": clusterName,
	}

	log.Printf("labels: %v", labels)
	log.Printf("get entry in setInstanceLabel")
	err = SetGkeLabel(name, resourceLocation, labels, labelFingerprint)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("The inserted instance %s has been  labeled successfully")

	return nil

}
