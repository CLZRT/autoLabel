package autolabel

import (
	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"context"
	"fmt"
)

func labelOperation(instanceId string, labels map[string]string) error {
	instance, err := getInstance(labels["project_id"], labels["zone"], labels["instance_id"])
	if err != nil {
		return err
	}
	labelFingerprint := instance.GetLabelFingerprint()
	err = setInstanceLabel(labels["project_id"], labels["zone"], labels["instance_id"], labels, &labelFingerprint)
	if err != nil {
		return err
	}
	return nil

}

// getInstance prints a name of a VM instance in the given zone in the specified project.
func getInstance(projectID, zone, instanceName string) (*computepb.Instance, error) {
	// projectID := "your_project_id"
	// zone := "europe-central2-b"
	// instanceName := "your_instance_name"

	ctx := context.Background()
	instancesClient, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("instancesClient: %w", err)
	}
	defer instancesClient.Close()
	reqInstance := &computepb.GetInstanceRequest{
		Project:  projectID,
		Zone:     zone,
		Instance: instanceName,
	}

	instance, err := instancesClient.Get(ctx, reqInstance)
	if err != nil {
		return nil, fmt.Errorf("instancesGet: %w", err)
	}

	return instance, nil
}

func setInstanceLabel(projectID, zone, instanceName string, labels map[string]string, labelFingerprint *string) error {
	// Add the creator label to the instance
	_, err := client.SetLabels(context.Background(), &computepb.SetLabelsInstanceRequest{
		Project:  projectID,
		Zone:     zone,
		Instance: instanceName,
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
