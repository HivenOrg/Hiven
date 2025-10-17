package storage

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Storage struct {
	client     *s3.Client
	bucketName string
}

func New(bucketName, region, awsAccessKeyID, awsSecretAccessKey, stage string, ctx context.Context) (*Storage, error) {

	var loadOptions []func(*config.LoadOptions) error
	loadOptions = append(loadOptions, config.WithRegion(region))

	// dev/test needs access keys explicitily defined
	// In prod IAM handles this
	if stage == "dev" || stage == "test" {
		loadOptions = append(loadOptions,
			config.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(awsAccessKeyID, awsSecretAccessKey, ""),
			),
		)
	}

	awsCfg, err := config.LoadDefaultConfig(ctx, loadOptions...)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS configuration: %w", err)
	}

	client := s3.NewFromConfig(awsCfg)

	return &Storage{client: client, bucketName: bucketName}, nil
}
