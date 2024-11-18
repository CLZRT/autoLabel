package dataproc

import (
	dataproc "cloud.google.com/go/dataproc/apiv1"
	"cloud.google.com/go/dataproc/apiv1/dataprocpb"
	"context"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"log"
)

func getCluster(clusterLocation map[string]string) (*dataprocpb.Cluster, error) {
	ctx := context.Background()
	endpoint := clusterLocation["region"] + "-dataproc.googleapis.com:443"
	client, err := dataproc.NewClusterControllerClient(ctx, option.WithEndpoint(endpoint))
	if err != nil {
		return nil, err
	}
	defer client.Close()
	getRequest := dataprocpb.GetClusterRequest{
		ProjectId:   clusterLocation["project_id"],
		Region:      clusterLocation["region"],
		ClusterName: clusterLocation["cluster_name"],
	}
	log.Println("Get Dataproc Cluster")
	cluster, err := client.GetCluster(ctx, &getRequest)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

func setClusterLabel(cluster *dataprocpb.Cluster, clusterLocation, labels map[string]string) error {
	ctx := context.Background()
	log.Println("New cluster client")
	endpoint := clusterLocation["region"] + "-dataproc.googleapis.com:443"
	client, err := dataproc.NewClusterControllerClient(ctx, option.WithEndpoint(endpoint))
	if err != nil {
		return err
	}
	defer client.Close()
	cluster.Labels = labels
	updateClusterRequest := dataprocpb.UpdateClusterRequest{
		ProjectId:   cluster.GetProjectId(),
		Region:      clusterLocation["region"],
		ClusterName: cluster.GetClusterName(),
		Cluster:     cluster,
		UpdateMask:  &fieldmaskpb.FieldMask{Paths: []string{"labels"}},
	}
	log.Println("Update Dataproc Cluster")
	_, err = client.UpdateCluster(ctx, &updateClusterRequest)
	if err != nil {
		return err
	}
	return nil
}

func getJob(jobLocation map[string]string) (*dataprocpb.Job, error) {
	ctx := context.Background()
	log.Println("New cluster client")
	endpoint := jobLocation["region"] + "-dataproc.googleapis.com:443"
	client, err := dataproc.NewJobControllerClient(ctx, option.WithEndpoint(endpoint))
	if err != nil {
		return nil, err
	}
	defer client.Close()
	log.Println("Get Dataproc Job")
	for i, v := range jobLocation {
		log.Println("Get Job", i, v)
	}
	getRequest := dataprocpb.GetJobRequest{

		// Required. The ID of the Google Cloud Platform project that the job
		// belongs to.
		ProjectId: jobLocation["project_id"],

		// Required. The Dataproc region in which to handle the request.
		Region: jobLocation["region"],

		// Required. The job ID.
		JobId: jobLocation["job_id"],
	}
	job, err := client.GetJob(ctx, &getRequest)
	if err != nil {
		return nil, err
	}
	return job, nil
}

func setJobLabel(job *dataprocpb.Job, jobLocation map[string]string, labels map[string]string) error {
	ctx := context.Background()
	log.Println("New job client")
	endpoint := jobLocation["region"] + "-dataproc.googleapis.com:443"
	client, err := dataproc.NewJobControllerClient(ctx, option.WithEndpoint(endpoint))
	if err != nil {
		return err
	}
	defer client.Close()
	job.Labels = labels

	updateJobRequest := dataprocpb.UpdateJobRequest{
		ProjectId: jobLocation["project_id"],
		Region:    jobLocation["region"],
		JobId:     jobLocation["job_id"],
		Job:       job,
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"labels"},
		},
	}
	log.Println("Update Dataproc Job")
	_, err = client.UpdateJob(ctx, &updateJobRequest)
	if err != nil {
		return err
	}
	return nil
}
