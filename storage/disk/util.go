package disk

import (
	"autolabel/logstruct"
	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"context"
	"fmt"
)

func getDisk(resourceLabels *logstruct.AuditResourceLabels) (*computepb.Disk, error) {
	ctx := context.Background()
	diskClient, err := compute.NewDisksRESTClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("diskClient: %w", err)
	}
	defer diskClient.Close()

	reqDisk := &computepb.GetDiskRequest{
		Project: resourceLabels.ProjectId,
		Zone:    resourceLabels.Zone,
		Disk:    resourceLabels.ResourceId,
	}
	disk, err := diskClient.Get(ctx, reqDisk)
	if err != nil {
		return nil, fmt.Errorf("disk: %w", err)
	}
	return disk, nil

}

func setDiskLabel(
	resourceLabels *logstruct.AuditResourceLabels,
	labels map[string]string, labelFingerprint *string) error {

	ctx := context.Background()
	diskClient, err := compute.NewDisksRESTClient(ctx)
	if err != nil {
		return fmt.Errorf("diskClient: %w", err)
	}
	defer diskClient.Close()

	reqDisk := &computepb.SetLabelsDiskRequest{
		Project:  resourceLabels.ProjectId,
		Zone:     resourceLabels.Zone,
		Resource: resourceLabels.ResourceId,
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
