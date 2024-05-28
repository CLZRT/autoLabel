package bigquery

import (
	"autolabel/logstruct"
	"cloud.google.com/go/bigquery"
	"context"
)

func GetDataset(resource *logstruct.AuditResourceLabels) (client *bigquery.Dataset, err error) {
	ctx := context.Background()
	newClient, err := bigquery.NewClient(ctx, resource.ProjectId)
	if err != nil {
		return nil, err
	}
	defer newClient.Close()

	dataset := newClient.Dataset(resource.ResourceId)
	return dataset, nil
}

func SetDatasetLabel(resource *logstruct.AuditResourceLabels, dataset *bigquery.Dataset, labels map[string]string) error {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, resource.ProjectId)
	if err != nil {
		return err
	}
	defer client.Close()

	metadata, err := dataset.Metadata(ctx)
	if err != nil {
		return err
	}
	metadataToUpdate := bigquery.DatasetMetadataToUpdate{}
	for key, value := range labels {
		metadataToUpdate.SetLabel(key, value)
	}
	// Todo dataset SetLabel
	_, err = dataset.Update(ctx, metadataToUpdate, metadata.ETag)
	if err != nil {
		return err
	}
	return nil
}
