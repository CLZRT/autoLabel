package dataproc

import (
	dataproc "cloud.google.com/go/dataproc/apiv1"
	"cloud.google.com/go/dataproc/apiv1/dataprocpb"
	"context"
	"log"
)

func getCluster() *dataprocpb.Cluster {
	ctx := context.Background()
	client, err := dataproc.NewClusterControllerRESTClient(ctx)
	if err != nil {
		return nil
	}
	cluster, err := client.GetCluster(ctx, &dataprocpb.GetClusterRequest{})
	if err != nil {
		return nil
	}
	return cluster

}
func setLabelCluster(cluster *dataprocpb.Cluster, labels map[string]string) {
	ctx := context.Background()
	client, err := dataproc.NewClusterControllerRESTClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create cluster controller REST client: %v", err)
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
	operation, err := client.UpdateCluster(ctx, &updateRequest)
	if err != nil {
		return
	}

}
