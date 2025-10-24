package storage

import (
	"bytes"
	"context"
	"fmt"
	"mime"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
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

func InferMimeTypeFromFilename(filename string) string {
	filename = strings.TrimSpace(filename)
	if filename == "" {
		return "application/octet-stream"
	}

	ext := strings.ToLower(filepath.Ext(filename))
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}

	return mimeType
}

func (storage Storage) Upload(ctx context.Context, key, mimeType string, data []byte) error {
	_, err := storage.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(storage.bucketName),
		Key:         aws.String(key),
		Body:        bytes.NewReader(data),
		ContentType: aws.String(mimeType),
	})
	if err != nil {
		return fmt.Errorf("failed to upload to S3: %w", err)
	}
	return nil
}

func (storage Storage) DownloadPresignedURL(ctx context.Context, key string, expiresInMinutes int) (string, error) {

	if strings.TrimSpace(key) == "" {
		return "", fmt.Errorf("invalid S3 key")
	}

	psClient := s3.NewPresignClient(storage.client)

	req, err := psClient.PresignGetObject(
		ctx,
		&s3.GetObjectInput{
			Bucket: aws.String(storage.bucketName),
			Key:    aws.String(key),
		},
		s3.WithPresignExpires(time.Duration(expiresInMinutes)*time.Minute),
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %w", err)
	}

	return req.URL, nil
}
