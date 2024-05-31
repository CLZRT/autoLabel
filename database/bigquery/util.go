package bigquery

import (
	"cloud.google.com/go/bigquery"
	"context"
	"errors"
	"google.golang.org/api/iterator"
	"log"
)

func GetDataset(resourceLocation map[string]string) (client *bigquery.Dataset, err error) {
	ctx := context.Background()
	newClient, err := bigquery.NewClient(ctx, resourceLocation["project-id"])
	if err != nil {
		return nil, err
	}
	defer newClient.Close()
	dataset := newClient.Dataset(resourceLocation["dataset-id"])
	return dataset, nil
}
func GetTable(resourceLocation map[string]string) (client *bigquery.Table, err error) {
	ctx := context.Background()
	newClient, err := bigquery.NewClient(ctx, resourceLocation["project-id"])
	if err != nil {
		return nil, err
	}
	defer newClient.Close()
	table := newClient.Dataset(resourceLocation["dataset-id"]).Table(resourceLocation["table-id"])
	return table, nil
}

func SetDatasetandTable(labels map[string]string, dataset *bigquery.Dataset) error {
	ctx := context.Background()
	// Set DataSet Label
	log.Printf("Get Entry into SetDatasetLabel")
	err := SetDatasetLabel(labels, dataset)
	if err != nil {
		return err
	}

	// Set Tabel Label belong to DataSet
	tableIterator := dataset.Tables(ctx)
	for {
		table, err := tableIterator.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		tableLabel := map[string]string{
			"created-by": labels["created-by"],
			"table-id":   table.TableID,
		}
		for k, v := range tableLabel {
			log.Printf("Set Table label %s to %s", k, v)
		}
		log.Printf("Get Entry into SetDatasetLabel")
		err = SetTableLabel(tableLabel, table)
		if err != nil {
			return err
		}
	}
	return nil
}
func SetDatasetLabel(labels map[string]string, dataset *bigquery.Dataset) error {
	ctx := context.Background()

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
func SetTableLabel(labels map[string]string, table *bigquery.Table) error {
	ctx := context.Background()

	tabelToUpdate := bigquery.TableMetadataToUpdate{}

	for key, value := range labels {
		tabelToUpdate.SetLabel(key, value)
	}
	tableMetadata, err := table.Metadata(ctx)
	if err != nil {
		return err
	}
	_, err = table.Update(ctx, tabelToUpdate, tableMetadata.ETag)
	if err != nil {
		return err
	}
	return nil
}
