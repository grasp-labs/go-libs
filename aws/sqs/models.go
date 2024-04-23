package sqs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"

	"github.com/grasp-labs/go-libs/config"
)

type APISqs interface {
	GetQueueUrl(ctx context.Context, params *sqs.GetQueueUrlInput, optFns ...func(*sqs.Options)) (*sqs.GetQueueUrlOutput, error)
	SendMessage(ctx context.Context, params *sqs.SendMessageInput, optFns ...func(*sqs.Options)) (*sqs.SendMessageOutput, error)
}

type ClientSqs interface {
	SendMsg(ctx context.Context, input map[string]types.MessageAttributeValue) error
}

type Client struct {
	api       APISqs
	queueName string
	queueURL  string
}

func NewClient(ctx context.Context, queueName string) (*Client, error) {
	cfg, err := config.NewConfig(ctx)
	if err != nil {
		return nil, err
	}

	c := &Client{api: sqs.NewFromConfig(cfg), queueName: queueName}
	if err := c.setQueueURL(ctx); err != nil {
		return nil, err
	}

	return c, nil
}

func NewClientWithAPI(ctx context.Context, api APISqs, queueName string) (*Client, error) {

	c := &Client{api: api, queueName: queueName}
	if err := c.setQueueURL(ctx); err != nil {
		return nil, err
	}

	return c, nil
}
