package gke

import (
	"autolabel/logstruct"
	container "cloud.google.com/go/container/apiv1"
	"cloud.google.com/go/container/apiv1/containerpb"
	"context"
)

func getCluster(resourceLabel *logstruct.AuditResourceLabels) (*containerpb.Cluster, error) {
	ctx := context.Background()
	client, err := container.NewClusterManagerClient(ctx)
	if err != nil {
		return nil, err
	}
	getRequest := containerpb.GetClusterRequest{
		ProjectId: resourceLabel.ProjectId,
		Zone:      resourceLabel.Zone,
		ClusterId: resourceLabel.ResourceId,
	}
	cluster, err := client.GetCluster(ctx, &getRequest)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

func setLabel(resourceLabel *logstruct.AuditResourceLabels, cluster *containerpb.Cluster, labels map[string]string) error {
	ctx := context.Background()
	client, err := container.NewClusterManagerClient(ctx)
	if err != nil {
		return err
	}
	setLabelRequest := containerpb.SetLabelsRequest{
		ProjectId:        resourceLabel.ProjectId,
		Zone:             cluster.Zone,
		ClusterId:        cluster.Id,
		ResourceLabels:   labels,
		LabelFingerprint: "",
		Name:             "",
	}
	_, err = client.SetLabels(ctx, &setLabelRequest)
	if err != nil {
		return err
	}
	return nil
}
