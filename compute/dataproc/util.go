package dataproc

import (
	"autolabel/logstruct"
	dataproc "cloud.google.com/go/dataproc/apiv1"
	"cloud.google.com/go/dataproc/apiv1/dataprocpb"
	"context"
	"log"
)

func GetCluster(resourceLabels *logstruct.AuditResourceLabels) (*dataprocpb.Cluster, error) {
	ctx := context.Background()
	client, err := dataproc.NewClusterControllerRESTClient(ctx)
	if err != nil {
		return nil, err
	}
	cluster, err := client.GetCluster(ctx, &dataprocpb.GetClusterRequest{
		ProjectId:   resourceLabels.ProjectId,
		Region:      resourceLabels.Zone,
		ClusterName: resourceLabels.ResourceId,
	})
	if err != nil {
		return nil, err
	}
	return cluster, nil

}
func SetLabelCluster(cluster *dataprocpb.Cluster, labels map[string]string) error {
	ctx := context.Background()
	client, err := dataproc.NewClusterControllerRESTClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create cluster controller REST client: %v", err)
		return err
	}
	updateRequest := dataprocpb.UpdateClusterRequest{
		Cluster: &dataprocpb.Cluster{
			ProjectId:            cluster.ProjectId,
			ClusterName:          cluster.ClusterName,
			Config:               cluster.Config,
			VirtualClusterConfig: cluster.VirtualClusterConfig,
			Labels:               labels,
			Status:               cluster.Status,
			StatusHistory:        cluster.StatusHistory,
			ClusterUuid:          cluster.ClusterUuid,
			Metrics:              cluster.Metrics,
		},
	}
	_, err = client.UpdateCluster(ctx, &updateRequest)
	if err != nil {
		return err
	}
	return nil
}
