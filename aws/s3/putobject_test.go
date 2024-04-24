package s3

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/grasp-labs/go-libs/mocks"
	"github.com/stretchr/testify/assert"
)

func TestClient_PutObject(t *testing.T) {
	type fields struct {
		api APIS3
	}
	type args struct {
		ctx    context.Context
		bucket string
		key    string
		data   []byte
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErrMsg string
		setup      func(f *fields)
	}{
		{
			name:   "ShouldErrorOnPutObject",
			fields: fields{},
			args: args{
				ctx:    context.Background(),
				bucket: "foo_bucket/",
				key:    "foo_key/",
				data:   []byte("foo_data"),
			},
			wantErrMsg: "foo_error",
			setup: func(f *fields) {
				s3Mock := mocks.NewAPIS3(t)
				s3Mock.
					EXPECT().
					PutObject(context.Background(), &s3.PutObjectInput{
						Bucket: aws.String("foo_bucket/"),
						Key:    aws.String("foo_key/"),
						Body:   bytes.NewReader([]byte("foo_data")),
					}).
					Return(nil, fmt.Errorf("foo_error")).
					Once()

				f.api = s3Mock
			},
		},
		{
			name:   "ShouldPutObject",
			fields: fields{},
			args: args{
				ctx:    context.Background(),
				bucket: "foo_bucket/",
				key:    "foo_key/",
				data:   []byte("foo_data"),
			},
			setup: func(f *fields) {
				s3Mock := mocks.NewAPIS3(t)
				s3Mock.
					EXPECT().
					PutObject(context.Background(), &s3.PutObjectInput{
						Bucket: aws.String("foo_bucket/"),
						Key:    aws.String("foo_key/"),
						Body:   bytes.NewReader([]byte("foo_data")),
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
			err := c.PutObject(tt.args.ctx, tt.args.bucket, tt.args.key, tt.args.data)
			if err != nil {
				assert.EqualError(t, err, tt.wantErrMsg)
				return
			}
		})
	}
}
