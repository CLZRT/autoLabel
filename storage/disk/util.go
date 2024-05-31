package disk

import (
	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"context"
	"fmt"
)

func GetDisk(resourceLocation map[string]string) (*computepb.Disk, error) {
	ctx := context.Background()
	diskClient, err := compute.NewDisksRESTClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("diskClient: %w", err)
	}
	defer diskClient.Close()

	reqDisk := &computepb.GetDiskRequest{
		Project: resourceLocation["project-id"],
		Zone:    resourceLocation["zone"],
		Disk:    resourceLocation["name"],
	}
	disk, err := diskClient.Get(ctx, reqDisk)
	if err != nil {
		return nil, fmt.Errorf("disk: %w", err)
	}
	return disk, nil

}

func SetDiskLabel(resourceLocation, labels map[string]string, labelFingerprint *string) error {

	ctx := context.Background()
	diskClient, err := compute.NewDisksRESTClient(ctx)
	if err != nil {
		return fmt.Errorf("diskClient: %w", err)
	}
	defer diskClient.Close()

	reqDisk := &computepb.SetLabelsDiskRequest{
		Project:  resourceLocation["project-id"],
		Zone:     resourceLocation["zone"],
		Resource: resourceLocation["name"],
		ZoneSetLabelsRequestResource: &computepb.ZoneSetLabelsRequest{
			Labels:           labels,
			LabelFingerprint: labelFingerprint,
		},
	}

	_, err = diskClient.SetLabels(ctx, reqDisk)
	if err != nil {
		return err
	}
	return nil
}
