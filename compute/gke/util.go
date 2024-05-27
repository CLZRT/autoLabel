package gke

import (
	"autolabel/logstruct"
	container "cloud.google.com/go/container/apiv1"
	"cloud.google.com/go/container/apiv1/containerpb"
	"context"
)

func GetCluster(protoPayload *logstruct.AuditLogProtoPayload) (*containerpb.Cluster, error) {
	ctx := context.Background()
	client, err := container.NewClusterManagerClient(ctx)
	if err != nil {
		return nil, err
	}
	getRequest := containerpb.GetClusterRequest{
		// Specified in the format `projects/*/locations/*/clusters/*`.
		Name: protoPayload.ResourceName,
	}
	cluster, err := client.GetCluster(ctx, &getRequest)
	if err != nil {
		return nil, err
	}
	return cluster, nil
}

func SetLabel(resourceName string, cluster *containerpb.Cluster, labels map[string]string) error {
	ctx := context.Background()
	client, err := container.NewClusterManagerClient(ctx)
	if err != nil {
		return err
	}
	setLabelRequest := containerpb.SetLabelsRequest{
		ResourceLabels:   labels,
		LabelFingerprint: cluster.GetLabelFingerprint(),
		Name:             resourceName,
	}
	_, err = client.SetLabels(ctx, &setLabelRequest)
	if err != nil {
		return err
	}
	return nil
}
