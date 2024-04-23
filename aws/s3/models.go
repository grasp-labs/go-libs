package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/grasp-labs/go-libs/config"
)

type APIS3 interface {
	GetObject(context.Context, *s3.GetObjectInput, ...func(*s3.Options)) (*s3.GetObjectOutput, error)
}

type ClientS3 interface {
	GetObject(ctx context.Context, bucket, key string) ([]byte, error)
}

type Client struct {
	api APIS3
}

func NewClient(c context.Context) (*Client, error) {
	cfg, err := config.NewConfig(c)
	if err != nil {
		return nil, err
	}

	return &Client{api: s3.NewFromConfig(cfg)}, nil
}
