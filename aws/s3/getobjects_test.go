package s3

import (
	"context"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/grasp-labs/go-libs/mocks"
	"github.com/stretchr/testify/assert"
)

const (
	bucket = "foo/"
	key    = "bar/"
	data   = "foo_bar"
)

func TestClient_GetObjects(t *testing.T) {
	type fields struct {
		api APIS3
	}
	type args struct {
		ctx    context.Context
		bucket string
		key    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    [][]byte
		wantErr string
		setup   func(f *fields)
	}{
		{
			name:   "ShouldFailOnListObjects",
			fields: fields{},
			args: args{
				ctx:    context.Background(),
				bucket: bucket,
				key:    key,
			},
			wantErr: "foo_error",
			setup: func(f *fields) {
				s3Mock := mocks.NewAPIS3(t)
				s3Mock.
					EXPECT().
					ListObjectsV2(context.Background(), &s3.ListObjectsV2Input{
						Bucket: aws.String(bucket),
						Prefix: aws.String(key),
					}).
					Return(nil, fmt.Errorf("foo_error")).
					Once()

				f.api = s3Mock
			},
		},
		{
			name:   "ShouldFailOnGetObject",
			fields: fields{},
			args: args{
				ctx:    context.Background(),
				bucket: bucket,
				key:    key,
			},
			wantErr: "foo_error",
			setup: func(f *fields) {
				s3Mock := mocks.NewAPIS3(t)
				s3Mock.
					EXPECT().
					ListObjectsV2(context.Background(), &s3.ListObjectsV2Input{
						Bucket: aws.String(bucket),
						Prefix: aws.String(key),
					}).
					Return(&s3.ListObjectsV2Output{Contents: []types.Object{
						{Key: aws.String(key)},
					}}, nil).
					Once()

				s3Mock.
					EXPECT().
					GetObject(context.Background(), &s3.GetObjectInput{
						Bucket: aws.String(bucket),
						Key:    aws.String(key),
					}).
					Return(nil, fmt.Errorf("foo_error")).
					Once()

				f.api = s3Mock
			},
		},
		{
			name:   "ShouldGetObjects",
			fields: fields{},
			args: args{
				ctx:    context.Background(),
				bucket: bucket,
				key:    key,
			},
			want: [][]byte{[]byte(data)},
			setup: func(f *fields) {
				s3Mock := mocks.NewAPIS3(t)
				s3Mock.
					EXPECT().
					ListObjectsV2(context.Background(), &s3.ListObjectsV2Input{
						Bucket: aws.String(bucket),
						Prefix: aws.String(key),
					}).
					Return(&s3.ListObjectsV2Output{Contents: []types.Object{
						{Key: aws.String(key)},
					}}, nil).
					Once()

				s3Mock.
					EXPECT().
					GetObject(context.Background(), &s3.GetObjectInput{
						Bucket: aws.String(bucket),
						Key:    aws.String(key),
					}).
					Return(&s3.GetObjectOutput{
						Body: io.NopCloser(strings.NewReader(data)),
					}, nil).
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
			got, err := c.GetObjects(tt.args.ctx, tt.args.bucket, tt.args.key)
			if err != nil {
				assert.EqualError(t, err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
