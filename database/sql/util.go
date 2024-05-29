package sql

import (
	"autolabel/logstruct"
	"context"
	"google.golang.org/api/sqladmin/v1"
)

func GetSql(resourceLabels *logstruct.AuditResourceLabels) (*sqladmin.DatabaseInstance, error) {

	ctx := context.Background()
	service, err := sqladmin.NewService(ctx)

	if err != nil {
		return nil, err
	}
	// - instance: Database instance ID. This does not include the project ID.
	// - project: Project ID of the project that contains the instance.
	instance, err := service.Instances.Get(resourceLabels.ProjectId, resourceLabels.InstanceId).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func SetSqlLabel(resource *logstruct.AuditResourceLabels, instance *sqladmin.DatabaseInstance, labels map[string]string) error {
	ctx := context.Background()
	service, err := sqladmin.NewService(ctx)
	if err != nil {
		return err
	}
	instance.Settings.UserLabels = labels

	_, err = service.Instances.Update(resource.ProjectId, resource.InstanceId, instance).Do()
	if err != nil {
		return err
	}
	return nil
}
