package autolabel

import (
	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"context"
	"fmt"
	"io"
	"os"
)

func labelOperation(instanceId string, labels map[string]string) error {
	writer := io.Writer(os.Stdout)
	instance, err := getInstance(writer, labels["project_id"], labels["zone"], labels["instance_id"])
	if err != nil {
		return err
	}
	labelFingerprint := instance.GetLabelFingerprint()
	err = setInstanceLabel(writer, labels["project_id"], labels["zone"], labels["instance_id"], labels, &labelFingerprint)
	if err != nil {
		return err
	}
	return nil

}

// getInstance prints a name of a VM instance in the given zone in the specified project.
func getInstance(w io.Writer, projectID, zone, instanceName string) (*computepb.Instance, error) {
	// projectID := "your_project_id"
	// zone := "europe-central2-b"
	// instanceName := "your_instance_name"

	ctx := context.Background()
	instancesClient, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("NewInstancesRESTClient: %w", err)
	}
	defer instancesClient.Close()

	reqInstance := &computepb.GetInstanceRequest{
		Project:  projectID,
		Zone:     zone,
		Instance: instanceName,
	}

	instance, err := instancesClient.Get(ctx, reqInstance)
	if err != nil {
		return nil, fmt.Errorf("unable to get instance: %w", err)
	}

	fmt.Fprintf(w, "Instance: %s\n", instance.GetName())

	return instance, nil
}

func setInstanceLabel(w io.Writer, projectID, zone, instanceName string, labels map[string]string, labelFingerprint *string) error {
	ctx := context.Background()
	instancesClient, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		return fmt.Errorf("NewInstancesRESTClient: %w", err)
	}
	defer instancesClient.Close()
	reqInstance := &computepb.SetLabelsInstanceRequest{
		Project:  projectID,
		Zone:     zone,
		Instance: instanceName,
		InstancesSetLabelsRequestResource: &computepb.InstancesSetLabelsRequest{
			LabelFingerprint: labelFingerprint,
			Labels:           labels,
		},
	}

	setLabels, err := instancesClient.SetLabels(ctx, reqInstance)
	if err != nil {
		return fmt.Errorf("unable to get instance: %w", err)
	}

	fmt.Fprintf(w, "Instance: %s\n", setLabels.Name())

	return nil

}
