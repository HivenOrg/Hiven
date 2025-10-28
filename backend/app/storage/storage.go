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

func (storage Storage) ListObjects(ctx context.Context, prefix *string) ([]string, error) {

	var keys []string

	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(storage.bucketName),
	}

	// Apply prefix if provided and not empty
	if prefix != nil && strings.TrimSpace(*prefix) != "" {
		input.Prefix = aws.String(strings.TrimSpace(*prefix))
	}

	paginator := s3.NewListObjectsV2Paginator(storage.client, input)

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to list objects: %w", err)
		}

		for _, obj := range page.Contents {
			if obj.Key != nil {
				keys = append(keys, *obj.Key)
			}
		}
	}

	return keys, nil
}

func (storage Storage) DeleteObject(ctx context.Context, key string) error {
	key = strings.TrimSpace(key)

	if key == "" {
		return fmt.Errorf("invalid S3 key: cannot be empty")
	}

	if strings.HasSuffix(key, "/") {
		return fmt.Errorf("invalid S3 key: cannot end with '/' (prefix detected)")
	}

	_, err := storage.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(storage.bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("failed to delete object %q: %w", key, err)
	}

	return nil
}
