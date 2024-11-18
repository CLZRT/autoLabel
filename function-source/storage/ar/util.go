package ar

import (
	artifactregistry "cloud.google.com/go/artifactregistry/apiv1beta2"
	artifactregistrypb "cloud.google.com/go/artifactregistry/apiv1beta2/artifactregistrypb"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func GetRepository(repoName string) (*artifactregistrypb.Repository, error) {
	// 获取当前的repository信息
	ctx := context.Background()
	client, err := artifactregistry.NewClient(ctx)
	getReq := &artifactregistrypb.GetRepositoryRequest{
		Name: repoName,
	}
	repo, err := client.GetRepository(ctx, getReq)
	if err != nil {
		return nil, err
	}

	return repo, err
}

func SetRepositoryLabel(repo *artifactregistrypb.Repository, labels map[string]string) error {
	// 设置标签
	ctx := context.Background()
	client, err := artifactregistry.NewClient(ctx)
	if repo.Labels == nil {
		repo.Labels = make(map[string]string)
	}
	for key, value := range labels {
		repo.Labels[key] = value
	}

	// 更新repository
	updateReq := &artifactregistrypb.UpdateRepositoryRequest{
		Repository: repo,
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"labels"},
		},
	}
	_, err = client.UpdateRepository(ctx, updateReq)
	if err != nil {
		return fmt.Errorf("SetLabels: %w", err)
	}
	return nil
}
