package s3

import (
	"context"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/grasp-labs/go-libs/mocks"
	"github.com/stretchr/testify/assert"
)

func TestClient_GetObject(t *testing.T) {
	type fields struct {
		api APIS3
	}
	type args struct {
		ctx    context.Context
		bucket string
		key    string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       []byte
		wantErrMsg string
		setup      func(f *fields)
	}{
		{
			name:   "ShouldErrorOnGetItem",
			fields: fields{},
			args: args{
				ctx:    context.Background(),
				bucket: "foo_bucket",
				key:    "foo_key",
			},
			wantErrMsg: "foo error",
			setup: func(f *fields) {
				s3Mock := mocks.NewAPIS3(t)
				s3Mock.EXPECT().
					GetObject(context.Background(), &s3.GetObjectInput{
						Bucket: aws.String("foo_bucket"),
						Key:    aws.String("foo_key"),
					}).
					Return(nil, fmt.Errorf("foo error")).
					Once()

				f.api = s3Mock
			},
		},
		{
			name:   "ShouldGetItem",
			fields: fields{},
			args: args{
				ctx:    context.Background(),
				bucket: "foo_bucket",
				key:    "foo_key",
			},
			want: []byte("foo_bar_body"),
			setup: func(f *fields) {
				s3Mock := mocks.NewAPIS3(t)
				s3Mock.EXPECT().
					GetObject(context.Background(), &s3.GetObjectInput{
						Bucket: aws.String("foo_bucket"),
						Key:    aws.String("foo_key"),
					}).
					Return(&s3.GetObjectOutput{
						Body: io.NopCloser(strings.NewReader("foo_bar_body")),
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
			got, err := c.GetObject(tt.args.ctx, tt.args.bucket, tt.args.key)
			if err != nil {
				assert.EqualError(t, err, tt.wantErrMsg)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
