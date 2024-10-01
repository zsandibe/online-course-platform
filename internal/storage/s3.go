package storage

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/zsandibe/online-course-platform/config"
)

func NewS3Client(ctx context.Context, cfg *config.Config) (*s3.Client, error) {

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		if service == s3.ServiceID && region == cfg.S3.SigningRegion {
			return aws.Endpoint{
				PartitionID:   cfg.S3.PartitionId,
				URL:           cfg.S3.Url,
				SigningRegion: cfg.S3.SigningRegion,
			}, nil
		}
		return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
	})

	// the method get data from environment variables
	awsCfg, err := awsConfig.LoadDefaultConfig(ctx, awsConfig.WithEndpointResolverWithOptions(customResolver))
	if err != nil {
		return nil, err
	}

	return s3.NewFromConfig(awsCfg), nil
}
