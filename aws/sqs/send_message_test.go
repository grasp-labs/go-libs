package sqs

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/stretchr/testify/assert"

	"github.com/grasp-labs/go-libs/mocks"
)

var input = map[string]types.MessageAttributeValue{
	"product_id": {
		DataType:    aws.String("String"),
		StringValue: aws.String("foo_product"),
	},
	"tenant_id": {
		DataType:    aws.String("String"),
		StringValue: aws.String("foo_product"),
	},
	"memory_mb": {
		DataType:    aws.String("String"),
		StringValue: aws.String("foo_product"),
	},
	"start_timestamp": {
		DataType:    aws.String("String"),
		StringValue: aws.String("foo_product"),
	},
	"end_timestamp": {
		DataType:    aws.String("String"),
		StringValue: aws.String("foo_product"),
	},
	"workflow": {
		DataType:    aws.String("String"),
		StringValue: aws.String("foo_product"),
	},
}

func TestClient_SendMsg(t *testing.T) {
	type fields struct {
		api       APISqs
		queueName string
		queueURL  string
	}
	type args struct {
		ctx   context.Context
		input map[string]types.MessageAttributeValue
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErrMsg string
		setup      func(f *fields)
	}{
		{
			name: "ShouldErrorOnSendMessage",
			fields: fields{
				queueName: "foo_name",
				queueURL:  "foo_url",
			},
			args: args{
				ctx:   context.Background(),
				input: input,
			},
			setup: func(f *fields) {
				apiMock := mocks.NewAPISqs(t)
				apiMock.EXPECT().
					SendMessage(context.Background(), &sqs.SendMessageInput{MessageBody: aws.String("foo_name"), QueueUrl: aws.String("foo_url"), MessageAttributes: input}).
					Return(nil, fmt.Errorf("foo error")).
					Once()

				f.api = apiMock
			},
			wantErrMsg: "foo error",
		},
		{
			name: "ShouldSendMessage",
			fields: fields{
				queueName: "foo_name",
				queueURL:  "foo_url",
			},
			args: args{
				ctx:   context.Background(),
				input: input,
			},
			setup: func(f *fields) {
				apiMock := mocks.NewAPISqs(t)
				apiMock.EXPECT().
					SendMessage(context.Background(), &sqs.SendMessageInput{MessageBody: aws.String("foo_name"), QueueUrl: aws.String("foo_url"), MessageAttributes: input}).
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
				api:       tt.fields.api,
				queueName: tt.fields.queueName,
				queueURL:  tt.fields.queueURL,
			}

			err := c.SendMsg(tt.args.ctx, tt.args.input)
			if err != nil {
				assert.EqualError(t, err, tt.wantErrMsg)
				return
			}
		})
	}
}

func TestClient_setQueueURL(t *testing.T) {
	type fields struct {
		api       APISqs
		queueName string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErrMsg string
		wantUrl    string
		setup      func(f *fields)
	}{
		{
			name: "ShouldErrorSetQueue",
			fields: fields{
				queueName: "foo_name",
			},
			args: args{
				ctx: context.Background(),
			},
			setup: func(f *fields) {
				apiMock := mocks.NewAPISqs(t)
				apiMock.EXPECT().
					GetQueueUrl(context.Background(), &sqs.GetQueueUrlInput{
						QueueName: aws.String("foo_name"),
					}).
					Return(nil, fmt.Errorf("foo error")).
					Once()

				f.api = apiMock
			},
			wantErrMsg: "foo error",
		},
		{
			name: "ShouldErrorOnEmptyURL",
			fields: fields{
				queueName: "foo_name",
			},
			args: args{
				ctx: context.Background(),
			},
			setup: func(f *fields) {
				apiMock := mocks.NewAPISqs(t)
				apiMock.EXPECT().
					GetQueueUrl(context.Background(), &sqs.GetQueueUrlInput{
						QueueName: aws.String("foo_name"),
					}).
					Return(&sqs.GetQueueUrlOutput{}, nil).
					Once()

				f.api = apiMock
			},
			wantErrMsg: "cannot find queue url",
		},
		{
			name: "ShouldSetQueue",
			fields: fields{
				queueName: "foo_name",
			},
			args: args{
				ctx: context.Background(),
			},
			setup: func(f *fields) {
				apiMock := mocks.NewAPISqs(t)
				apiMock.EXPECT().
					GetQueueUrl(context.Background(), &sqs.GetQueueUrlInput{
						QueueName: aws.String("foo_name"),
					}).
					Return(&sqs.GetQueueUrlOutput{QueueUrl: aws.String("foo_url")}, nil).
					Once()

				f.api = apiMock
			},
			wantUrl: "foo_url",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(&tt.fields)

			c := &Client{
				api:       tt.fields.api,
				queueName: tt.fields.queueName,
			}
			if err := c.setQueueURL(context.Background()); err != nil {
				assert.EqualError(t, err, tt.wantErrMsg)
				return
			}
			assert.Equal(t, tt.wantUrl, c.queueURL)
		})
	}
}
