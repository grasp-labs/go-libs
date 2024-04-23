package config

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func NewConfig(ctx context.Context) (aws.Config, error) {
	var (
		cfg aws.Config
		err error
	)

	if os.Getenv(buildingMode) == testMode {
		cfg, err = config.LoadDefaultConfig(ctx,
			config.WithRegion(region),
			config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
				func(service, region string, options ...interface{}) (aws.Endpoint, error) {
					return aws.Endpoint{URL: localUrl}, nil
				}),
			),
		)
	} else {
		cfg, err = config.LoadDefaultConfig(ctx, config.WithRegion(region))
	}
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
