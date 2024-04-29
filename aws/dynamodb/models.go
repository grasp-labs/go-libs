package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/grasp-labs/go-libs/config"
)

type APIDynamoDB interface {
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
	Query(ctx context.Context, input *dynamodb.QueryInput, f ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)
}

type ClientDynamoDB interface {
	PutItem(ctx context.Context, table string, itemToPut any) error
	Query(ctx context.Context, tableName, key, value string) ([]map[string]any, error)
}

type Client struct {
	api APIDynamoDB
}

func NewClient(ctx context.Context) (*Client, error) {
	cfg, err := config.NewConfig(ctx)
	if err != nil {
		return nil, err
	}

	return &Client{api: dynamodb.NewFromConfig(cfg)}, nil
}

func NewClientWithAPI(api APIDynamoDB) (*Client, error) {
	return &Client{api: api}, nil
}
