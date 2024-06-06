package deploy

import (
	deploy "cloud.google.com/go/deploy/apiv1"
	"cloud.google.com/go/deploy/apiv1/deploypb"
	"context"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"log"
)

func getTarget(targetName string) (*deploypb.Target, error) {
	ctx := context.Background()
	client, err := deploy.NewCloudDeployClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	log.Println("get target: " + targetName)
	target, err := client.GetTarget(ctx, &deploypb.GetTargetRequest{
		// Required. Name of the `Target`. Format must be
		// `projects/{project_id}/locations/{location_name}/targets/{target_name}`.
		Name: targetName,
	})
	if err != nil {
		return nil, err
	}
	return target, nil

}

func setTarget(labels map[string]string, target *deploypb.Target) error {
	ctx := context.Background()
	client, err := deploy.NewCloudDeployClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	target.Labels = labels
	log.Println("set target " + target.Name)
	_, err = client.UpdateTarget(ctx, &deploypb.UpdateTargetRequest{
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"labels"},
		},
		Target: target,
	})
	if err != nil {
		return err
	}

	return nil
}

func getTargetfromRollout(rolloutName string) (*deploypb.Target, error) {
	ctx := context.Background()
	client, err := deploy.NewCloudDeployClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	log.Println("get rollout: " + rolloutName)
	rollout, err := client.GetRollout(ctx, &deploypb.GetRolloutRequest{
		Name: rolloutName,
	})
	if err != nil {
		return nil, err
	}

	targetId := rollout.GetTargetId()
	targetLocation := "projects/testftl-4-6/locations/us-central1/targets/" + targetId
	log.Println("get target: " + targetLocation)
	target, err := getTarget(targetLocation)
	if err != nil {
		return nil, err
	}
	log.Println("get target: " + target.TargetId + "success")
	return target, nil
}
