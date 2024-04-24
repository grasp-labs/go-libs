package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (c *Client) DeleteObject(ctx context.Context, bucketName, key string) error {
	_, err := c.api.DeleteObject(ctx,
		&s3.DeleteObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(key),
		})
	if err != nil {
		return err
	}

	return nil
}
