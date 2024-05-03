package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// GetObjects retrieves objects from the specified bucket with keys prefixed by the given key.
// It returns a map where key is s3 bucket key, and value is s3 object data.
func (c *Client) GetObjects(ctx context.Context, bucket, key string) (map[string][]byte, error) {
	keys, err := c.listObjects(ctx, bucket, key)
	if err != nil {
		return nil, err
	}

	result := map[string][]byte{}
	for _, v := range keys {
		obj, err := c.GetObject(ctx, bucket, *v.Key)
		if err != nil {
			return nil, err
		}
		result[*v.Key] = obj
	}

	return result, nil
}

func (c *Client) listObjects(ctx context.Context, bucket, key string) ([]types.Object, error) {
	result, err := c.api.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(key),
	})
	if err != nil {
		return nil, err
	}

	return result.Contents, err
}
