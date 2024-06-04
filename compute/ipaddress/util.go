package ipaddress

import (
	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"context"
)

func GetIpaddress(ipLocation map[string]string) (*computepb.Address, error) {
	ctx := context.Background()
	addressService, err := compute.NewAddressesRESTClient(ctx)
	if err != nil {
		return nil, err
	}
	defer addressService.Close()
	address, err := addressService.Get(ctx, &computepb.GetAddressRequest{
		// Project ID for this request.
		Project: ipLocation["project_id"],
		// Name of the region for this request.
		Region: ipLocation["region"],
		// Name of the address resource to return.
		Address: ipLocation["name"],
	})
	if err != nil {
		return nil, err
	}
	return address, nil

}

func SetIpaddress(ipLocation, labels map[string]string, ipaddress *computepb.Address) error {
	ctx := context.Background()
	addressService, err := compute.NewAddressesRESTClient(ctx)
	if err != nil {
		return err
	}
	defer addressService.Close()

	_, err = addressService.SetLabels(ctx, &computepb.SetLabelsAddressRequest{
		Project: ipLocation["project_id"],
		Region:  ipLocation["region"],
		RegionSetLabelsRequestResource: &computepb.RegionSetLabelsRequest{
			LabelFingerprint: ipaddress.LabelFingerprint,
			Labels:           labels,
		},
		Resource: ipLocation["name"],
	})
	if err != nil {
		return err
	}
	return nil
}
