package autolabel

import (
	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"context"
	"fmt"
	"log"
)

func labelOperation(labels map[string]string) error {
	instance, err := getInstance(labels["project-id"], labels["zone"], labels["instance-name"])
	if err != nil {
		log.Printf("instance %s not exist", labels["instance-name"])
		return err
	}
	labelFingerprint := instance.GetLabelFingerprint()
	err = setInstanceLabel(labels, &labelFingerprint)
	if err != nil {
		log.Printf("label instance %s not exist", labels["instance-name"])
		return err
	}
	return nil

}

// getInstance prints a name of a VM instance in the given zone in the specified project.
func getInstance(projectID, zone, instanceName string) (*computepb.Instance, error) {

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

func setInstanceLabel(labels map[string]string, labelFingerprint *string) error {
	ctx := context.Background()
	instancesClient, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		return fmt.Errorf("instancesClient: %w", err)
	}
	defer instancesClient.Close()

	// Add the labels to the instance
	_, err = instancesClient.SetLabels(context.Background(), &computepb.SetLabelsInstanceRequest{
		Project:  labels["project-id"],
		Zone:     labels["zone"],
		Instance: labels["instance-name"],
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
