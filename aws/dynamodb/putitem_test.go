package dynamodb

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stretchr/testify/assert"

	"github.com/grasp-labs/go-libs/mocks"
)

type item struct {
	ID       string `json:"id" dynamodbav:"id"`
	TenantID string `json:"tenant_id" dynamodbav:"tenant_id"`
}

func TestClient_PutItem(t *testing.T) {
	type fields struct {
		api APIDynamoDB
	}
	type args struct {
		ctx       context.Context
		table     string
		itemToPut any
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErrMsg string
		setup      func(f *fields)
	}{
		{
			name: "ShouldErrorOnPutItem",
			args: args{
				ctx:   context.Background(),
				table: "dynamo-table",
				itemToPut: item{
					ID:       "foo-id",
					TenantID: "foo-tenant",
				},
			},
			wantErrMsg: "foo-dynamo",
			setup: func(f *fields) {
				apiMock := mocks.NewAPIDynamoDB(t)

				itemToPut, err := attributevalue.MarshalMap(item{
					ID:       "foo-id",
					TenantID: "foo-tenant",
				})
				if err != nil {
					assert.NoError(t, err)
					return
				}

				apiMock.EXPECT().
					PutItem(context.Background(), &dynamodb.PutItemInput{
						TableName: aws.String("dynamo-table"), Item: itemToPut,
					}).
					Return(nil, fmt.Errorf("foo-dynamo")).
					Once()

				f.api = apiMock
			},
		},
		{
			name: "ShouldPutItem",
			args: args{
				ctx:   context.Background(),
				table: "dynamo-table",
				itemToPut: item{
					ID:       "foo-id",
					TenantID: "foo-tenant",
				},
			},
			setup: func(f *fields) {
				apiMock := mocks.NewAPIDynamoDB(t)

				itemToPut, err := attributevalue.MarshalMap(item{
					ID:       "foo-id",
					TenantID: "foo-tenant",
				})
				if err != nil {
					assert.NoError(t, err)
					return
				}

				apiMock.EXPECT().
					PutItem(context.Background(), &dynamodb.PutItemInput{
						TableName: aws.String("dynamo-table"), Item: itemToPut,
					}).
					Return(nil, nil).
					Once()

				f.api = apiMock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(&tt.fields)

			c := &Client{
				api: tt.fields.api,
			}
			if err := c.PutItem(tt.args.ctx, tt.args.table, tt.args.itemToPut); err != nil {
				assert.EqualError(t, err, tt.wantErrMsg)
			}
		})
	}
}
