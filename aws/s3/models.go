package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/grasp-labs/go-libs/config"
)

type APIS3 interface {
	GetObject(ctx context.Context, input *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
	PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
	DeleteObject(ctx context.Context, params *s3.DeleteObjectInput, optFns ...func(*s3.Options)) (*s3.DeleteObjectOutput, error)
}

type ClientS3 interface {
	GetObject(ctx context.Context, bucket, key string) ([]byte, error)
	PutObject(ctx context.Context, bucket, key string, data []byte) error
	DeleteObject(ctx context.Context, bucketName, key string) error
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
