package s3

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/grasp-labs/go-libs/mocks"
	"github.com/stretchr/testify/assert"
)

func TestClient_DeleteObject(t *testing.T) {
	type fields struct {
		api APIS3
	}
	type args struct {
		ctx        context.Context
		bucketName string
		key        string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErrMsg string
		setup      func(f *fields)
	}{
		{
			name:   "ShouldErrorOnDeleteObject",
			fields: fields{},
			args: args{
				ctx:        context.Background(),
				bucketName: "foo_bucket/",
				key:        "foo_key/",
			},
			wantErrMsg: "foo_error",
			setup: func(f *fields) {
				s3Mock := mocks.NewAPIS3(t)
				s3Mock.
					EXPECT().
					DeleteObject(context.Background(), &s3.DeleteObjectInput{
						Bucket: aws.String("foo_bucket/"),
						Key:    aws.String("foo_key/"),
					}).
					Return(nil, fmt.Errorf("foo_error")).
					Once()

				f.api = s3Mock
			},
		},
		{
			name:   "ShouldDeleteObject",
			fields: fields{},
			args: args{
				ctx:        context.Background(),
				bucketName: "foo_bucket/",
				key:        "foo_key/",
			},
			setup: func(f *fields) {
				s3Mock := mocks.NewAPIS3(t)
				s3Mock.
					EXPECT().
					DeleteObject(context.Background(), &s3.DeleteObjectInput{
						Bucket: aws.String("foo_bucket/"),
						Key:    aws.String("foo_key/"),
					}).
					Return(nil, nil).
					Once()

				f.api = s3Mock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(&tt.fields)
			c := &Client{
				api: tt.fields.api,
			}

			err := c.DeleteObject(tt.args.ctx, tt.args.bucketName, tt.args.key)
			if err != nil {
				assert.EqualError(t, err, tt.wantErrMsg)
			}
		})
	}
}
