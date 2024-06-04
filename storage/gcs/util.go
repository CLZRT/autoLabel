package gcs

import (
	"context"
	"google.golang.org/api/storage/v1"
	"log"
)

func Getbucket(bucketName string) (*storage.Bucket, error) {
	ctx := context.Background()
	service, err := storage.NewService(ctx)
	if err != nil {
		return nil, err
	}
	bucket, err := service.Buckets.Get(bucketName).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	return bucket, nil
}

func SetBucketLabel(bucketName string, labels map[string]string, bucket *storage.Bucket) error {
	ctx := context.Background()
	service, err := storage.NewService(ctx)
	if err != nil {
		return err
	}
	bucket.Labels = labels
	log.Println("Set bucket: ", bucketName)
	_, err = service.Buckets.Patch(bucketName, bucket).Context(ctx).Do()
	if err != nil {
		return err
	}
	return nil
}
