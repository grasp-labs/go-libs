package paramstore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func (c *Client) GetParameter(ctx context.Context, name string, withDecryption bool) (*ssm.GetParameterOutput, error) {
	return c.api.GetParameter(ctx, &ssm.GetParameterInput{
		Name:           aws.String(name),
		WithDecryption: aws.Bool(withDecryption),
	})
}
