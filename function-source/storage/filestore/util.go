package filestore

import (
	filestore "cloud.google.com/go/filestore/apiv1"
	"cloud.google.com/go/filestore/apiv1/filestorepb"
	"context"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"log"
)

func GetInstance(resourceName string) (*filestorepb.Instance, error) {
	ctx := context.Background()
	client, err := filestore.NewCloudFilestoreManagerClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	log.Println("Get Filestore instance:", resourceName)
	instance, err := client.GetInstance(ctx, &filestorepb.GetInstanceRequest{
		// Required. The instance resource name, in the format
		// `projects/{project_id}/locations/{location}/instances/{instance_id}`.
		Name: resourceName,
	})
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func SetInstanceLabel(labels map[string]string, instance *filestorepb.Instance) error {
	ctx := context.Background()
	client, err := filestore.NewCloudFilestoreManagerClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	log.Println("Set Filestore instance:", instance.GetName())
	instance.Labels = labels
	_, err = client.UpdateInstance(ctx, &filestorepb.UpdateInstanceRequest{
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"labels"},
		},
		Instance: instance,
	})
	if err != nil {
		return err
	}
	return nil
}

func GetBackup(resourceName string) (*filestorepb.Backup, error) {
	ctx := context.Background()
	client, err := filestore.NewCloudFilestoreManagerClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	log.Println("Get Filestore backup:", resourceName)
	backup, err := client.GetBackup(ctx, &filestorepb.GetBackupRequest{
		// Required. The backup resource name, in the format
		// `projects/{project_number}/locations/{location}/backups/{backup_id}`.
		Name: resourceName,
	})
	if err != nil {
		return nil, err
	}
	return backup, nil
}

func SetBackupLabel(labels map[string]string, backup *filestorepb.Backup) error {
	ctx := context.Background()
	client, err := filestore.NewCloudFilestoreManagerClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()
	log.Println("Set Filestore backup:", backup)
	backup.Labels = labels
	_, err = client.UpdateBackup(ctx, &filestorepb.UpdateBackupRequest{
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"labels"},
		},
		Backup: backup,
	})
	if err != nil {
		return err
	}
	return nil
}
