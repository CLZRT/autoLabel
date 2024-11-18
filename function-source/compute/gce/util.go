package gce

import (
	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"context"
	"fmt"
)

// getInstance prints a name of a VM instance in the given zone in the specified project.
func getInstance(resourceLocation map[string]string) (*computepb.Instance, error) {

	ctx := context.Background()
	instancesClient, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("instancesClient: %w", err)
	}
	defer instancesClient.Close()
	reqInstance := &computepb.GetInstanceRequest{
		Project:  resourceLocation["project-id"],
		Zone:     resourceLocation["zone"],
		Instance: resourceLocation["instance-id"],
	}

	instance, err := instancesClient.Get(ctx, reqInstance)
	if err != nil {
		return nil, fmt.Errorf("instancesGet: %w", err)
	}

	return instance, nil
}

func setInstanceLabel(resourceLocation, labels map[string]string, labelFingerprint *string) error {
	ctx := context.Background()
	instancesClient, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		return fmt.Errorf("instancesClient: %w", err)
	}
	defer instancesClient.Close()

	// Add the labels to the instance
	_, err = instancesClient.SetLabels(context.Background(), &computepb.SetLabelsInstanceRequest{
		Project:  resourceLocation["project-id"],
		Zone:     resourceLocation["zone"],
		Instance: resourceLocation["instance-id"],
		InstancesSetLabelsRequestResource: &computepb.InstancesSetLabelsRequest{
			LabelFingerprint: labelFingerprint,
			Labels:           labels,
		},
	})
	if err != nil {
		return fmt.Errorf("SetLabels: %w", err)
	}
	return nil
}
