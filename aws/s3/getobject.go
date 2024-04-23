package s3

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// GetObject Retrieves an object from Amazon S3 based on bucket name and object key.
func (c *Client) GetObject(ctx context.Context, bucket, key string) ([]byte, error) {
	object, err := c.api.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}

	defer object.Body.Close()

	return io.ReadAll(object.Body)
}
