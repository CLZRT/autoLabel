package dataproc

import (
	"clzrt.io/autolabel/struct/logstruct"
	"fmt"
	"log"
	"regexp"
	"strings"
)

//	func DataProc(auditlog *logstruct.ClusterlogDP) error {
//		resourceName := auditlog.ProtoPayload.ResourceName
//		if strings.Contains(resourceName, "cluster") {
//			err := dataprocCluster(auditlog)
//			if err != nil {
//				return err
//			}
//			log.Println("set label in dataproc's cluster success")
//		} else if strings.Contains(resourceName, "job") {
//			err := dataprocJob(auditlog)
//			if err != nil {
//				return err
//			}
//			log.Println("set label in dataproc's job success")
//		}
//		//Todo: new error excluded
//		return nil
//	}
func Cluster(logAudit *logstruct.ClusterlogDP) error {

	resourceNameArray := strings.Split(logAudit.ProtoPayload.ResourceName, "/")
	clusterLocation := map[string]string{
		"project_id":   resourceNameArray[1],
		"region":       resourceNameArray[3],
		"cluster_name": resourceNameArray[5],
	}
	if resourceNameArray[5] == "" {
		return fmt.Errorf("resource name is invalid")
	}
	// Label Cluster
	cluster, err := getCluster(clusterLocation)
	if err != nil {
		return err
	}
	if cluster == nil {
		return fmt.Errorf("cluster is nil")
	}
	// break if cluster's state != running
	if cluster.Status.State.String() != "RUNNING" {
		log.Println(cluster.Status.String())
		return fmt.Errorf("this cluster is not running")
	}
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9-]+")
	creatorString := labelSanitizer.ReplaceAllString(strings.ToLower(logAudit.ProtoPayload.AuthenticationInfo.PrincipalEmail), "-")
	labels := map[string]string{
		"created-by":   creatorString,
		"cluster-id":   cluster.ClusterUuid,
		"cluster-name": cluster.ClusterName,
		"region":       clusterLocation["region"],
	}
	err = setClusterLabel(cluster, clusterLocation, labels)
	if err != nil {
		return err
	}
	log.Println("set label in dataproc's cluster success")
	return nil
	// Label GCE

	// Label GCS
}

func DataprocJob(logAudit *logstruct.JoblogDP) error {

	resourceNameArray := strings.Split(logAudit.ProtoPayload.ResourceName, "/")

	jobLocation := map[string]string{
		"project_id": resourceNameArray[1],
		"region":     resourceNameArray[3],
		"job_id":     resourceNameArray[5],
	}
	// Get job
	job, err := getJob(jobLocation)
	if err != nil {
		return err
	}
	// Set label in job
	labelSanitizer := regexp.MustCompile("[^a-zA-Z0-9-]+")
	creatorString := labelSanitizer.ReplaceAllString(strings.ToLower(logAudit.ProtoPayload.AuthenticationInfo.PrincipalEmail), "-")
	labels := map[string]string{
		"job-created-by": creatorString,
		"job-uuid":       job.GetJobUuid(),
		"job-name":       jobLocation["job_id"],
		"region":         jobLocation["region"],
	}
	err = setJobLabel(job, jobLocation, labels)
	if err != nil {
		return err
	}
	log.Println("set label in dataproc's job success")
	return nil
}
