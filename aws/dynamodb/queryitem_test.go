package dynamodb

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/grasp-labs/go-libs/mocks"
	"github.com/stretchr/testify/assert"
)

const (
	table    = "workflow_status_test"
	key      = "workflow_id"
	keyValue = "d3620960-d2cc-4419-930d-78a56b92a206"
)

func TestClient_Query(t *testing.T) {
	type fields struct {
		api APIDynamoDB
	}
	type args struct {
		ctx       context.Context
		tableName string
		key       string
		value     string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       []map[string]any
		wantErrMsg string
		setup      func(f *fields)
	}{
		{
			name:   "ShouldFailOnQueryItem",
			fields: fields{},
			args: args{
				ctx:       context.Background(),
				tableName: table,
				key:       key,
				value:     keyValue,
			},
			want:       []map[string]any{{key: keyValue}},
			wantErrMsg: "foo_error",
			setup: func(f *fields) {
				dbMock := mocks.NewAPIDynamoDB(t)

				keyEx := expression.Key(key).Equal(expression.Value(keyValue))
				expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
				if err != nil {
					assert.NoError(t, err)
					return
				}

				dbMock.
					EXPECT().
					Query(context.Background(), &dynamodb.QueryInput{
						TableName:                 aws.String(table),
						ExpressionAttributeNames:  expr.Names(),
						ExpressionAttributeValues: expr.Values(),
						KeyConditionExpression:    expr.KeyCondition()}).
					Return(nil, fmt.Errorf("foo_error")).
					Once()

				f.api = dbMock
			},
		},
		{
			name:   "ShouldQueryItem",
			fields: fields{},
			args: args{
				ctx:       context.Background(),
				tableName: table,
				key:       key,
				value:     keyValue,
			},
			want: []map[string]any{{key: keyValue}},
			setup: func(f *fields) {
				dbMock := mocks.NewAPIDynamoDB(t)

				keyEx := expression.Key(key).Equal(expression.Value(keyValue))
				expr, err := expression.NewBuilder().WithKeyCondition(keyEx).Build()
				if err != nil {
					assert.NoError(t, err)
					return
				}

				dbMock.
					EXPECT().
					Query(context.Background(), &dynamodb.QueryInput{
						TableName:                 aws.String(table),
						ExpressionAttributeNames:  expr.Names(),
						ExpressionAttributeValues: expr.Values(),
						KeyConditionExpression:    expr.KeyCondition()}).
					Return(&dynamodb.QueryOutput{Items: []map[string]types.AttributeValue{{key: &types.AttributeValueMemberS{Value: keyValue}}}}, nil).
					Once()

				f.api = dbMock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(&tt.fields)

			c := &Client{
				api: tt.fields.api,
			}
			got, err := c.Query(tt.args.ctx, tt.args.tableName, tt.args.key, tt.args.value)
			if err != nil {
				assert.EqualError(t, err, tt.wantErrMsg)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
