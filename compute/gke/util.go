package gke

import (
	container "cloud.google.com/go/container/apiv1"
	"cloud.google.com/go/container/apiv1/containerpb"
	"context"
	"fmt"
	"log"
)

// getInstance prints a name of a VM instance in the given zone in the specified project.
func GetGke(name string, resourceLocation map[string]string) (*containerpb.Cluster, error) {
	ctx := context.Background()
	c, err := container.NewClusterManagerClient(ctx)

	if err != nil {
		return nil, err
	}
	defer c.Close()
	reqGke := &containerpb.GetClusterRequest{
		ProjectId: resourceLocation["project-id"],
		Zone:      resourceLocation["zone"],
		ClusterId: resourceLocation["cluster-id"],
		Name:      name,
	}
	gke, err := c.GetCluster(ctx, reqGke)
	print("success")
	if err != nil {
		return nil, err
	}

	return gke, nil

}

func SetGkeLabel(name string, resourceLocation, labels map[string]string, labelFingerprint string) error {
	ctx := context.Background()
	c, err := container.NewClusterManagerClient(ctx)
	if err != nil {
		return err
	}
	defer c.Close()
	println("----------------")
	println(resourceLocation["project-id"])
	println(resourceLocation["zone"])
	println(resourceLocation["cluster-id"])
	println(name)
	for k, v := range labels {
		log.Println(k + "is" + v)
	}
	println(labelFingerprint)
	// Add the labels to the instance
	_, err = c.SetLabels(context.Background(), &containerpb.SetLabelsRequest{
		ProjectId:        resourceLocation["project-id"],
		Zone:             resourceLocation["zone"],
		ClusterId:        resourceLocation["cluster-id"],
		Name:             name,
		ResourceLabels:   labels,
		LabelFingerprint: labelFingerprint,
	})
	if err != nil {
		return fmt.Errorf("SetLabels: %w", err)
	}
	return nil
}
