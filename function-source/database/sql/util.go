package sql

import (
	"context"
	"fmt"
	"google.golang.org/api/sqladmin/v1"
	"log"
)

func GetDatabase(resourceLocation map[string]string) (*sqladmin.DatabaseInstance, error) {

	ctx := context.Background()
	service, err := sqladmin.NewService(ctx)
	if err != nil {
		return nil, err
	}
	// - instance: Database instance ID. This does not include the project ID.
	// - project: Project ID of the project that contains the instance.
	instance, err := service.Instances.Get(resourceLocation["project-id"], resourceLocation["database-name"]).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func SetDatabaseLabel(resourceLocation, labels map[string]string) error {
	ctx := context.Background()
	log.Printf("Now: New sql Service")
	service, err := sqladmin.NewService(ctx)
	if err != nil {
		return fmt.Errorf("DatabaseClient: %w", err)
	}
	log.Printf("Now: Get sql instance")

	instance, err := service.Instances.Get(resourceLocation["project-id"], resourceLocation["database-name"]).Context(ctx).Do()
	if err != nil {
		return err
	}

	instance.Settings.UserLabels = labels

	log.Printf("Now: Patch sql instance")
	_, err = service.Instances.Patch(resourceLocation["project-id"], resourceLocation["database-name"], instance).Context(ctx).Do()

	if err != nil {
		return err
	}
	return nil
}
