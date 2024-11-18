package memory

import (
	redis "cloud.google.com/go/redis/apiv1"
	"cloud.google.com/go/redis/apiv1/redispb"
	"context"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func GetMemoryStore(resourceName string) (*redispb.Instance, error) {
	ctx := context.Background()
	client, err := redis.NewCloudRedisRESTClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	getRequest := redispb.GetInstanceRequest{
		// `projects/{project_id}/locations/{location_id}/instances/{instance_id}`
		Name: resourceName,
	}
	instance, err := client.GetInstance(ctx, &getRequest)
	if err != nil {
		return nil, err
	}
	return instance, nil

}

func SetLabelMemoryStore(instance *redispb.Instance, labels map[string]string) error {
	ctx := context.Background()
	client, err := redis.NewCloudRedisRESTClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	instance.Labels = labels
	paths := []string{"labels"}
	updateRequest := redispb.UpdateInstanceRequest{
		//   - `displayName`
		//   - `labels`
		//   - `memorySizeGb`
		//   - `redisConfig`
		//   - `replica_count`
		UpdateMask: &fieldmaskpb.FieldMask{Paths: paths},
		Instance:   instance,
	}
	_, err = client.UpdateInstance(ctx, &updateRequest)
	if err != nil {
		return err
	}
	return nil

}
