package sqs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

func (c *Client) setQueueURL(ctx context.Context) error {
	result, err := c.api.GetQueueUrl(ctx, &sqs.GetQueueUrlInput{
		QueueName: aws.String(c.queueName),
	})
	if err != nil {
		return err
	}

	if result.QueueUrl == nil {
		return fmt.Errorf("cannot find queue url")
	}

	c.queueURL = *result.QueueUrl
	return nil
}

func (c *Client) SendMsg(ctx context.Context, input map[string]types.MessageAttributeValue) error {
	if _, err := c.api.SendMessage(ctx, &sqs.SendMessageInput{
		MessageBody:       aws.String(c.queueName),
		MessageAttributes: input,
		QueueUrl:          aws.String(c.queueURL),
	}); err != nil {
		return err
	}

	return nil
}
