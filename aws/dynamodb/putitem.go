package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

func (c *Client) PutItem(ctx context.Context, table string, itemToPut any) error {
	item, err := attributevalue.MarshalMap(itemToPut)
	if err != nil {
		return err
	}

	_, err = c.api.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(table), Item: item,
	})
	if err != nil {
		return err
	}

	return nil
}
