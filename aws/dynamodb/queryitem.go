package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (c *Client) Query(ctx context.Context, tableName, key, value string) ([]map[string]any, error) {
	keyEx := expression.Key(key).Equal(expression.Value(value))
	expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
	if err != nil {
		return nil, err
	}
	query, err := c.api.Query(ctx, &dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		KeyConditionExpression:    expr.KeyCondition()})
	if err != nil {
		return nil, err
	}

	var result []map[string]any
	if err := attributevalue.UnmarshalListOfMaps(query.Items, &result); err != nil {
		return nil, err
	}

	return result, nil
}
