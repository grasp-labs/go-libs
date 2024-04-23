package paramstore

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/stretchr/testify/assert"

	"github.com/grasp-labs/go-libs/mocks"
)

var (
	paramInput = &ssm.GetParameterInput{
		Name:           aws.String("foo_param"),
		WithDecryption: aws.Bool(true),
	}

	paramOutput = &ssm.GetParameterOutput{
		Parameter: &types.Parameter{
			Value: aws.String("foo value"),
		},
	}
)

func TestClient_GetParameter(t *testing.T) {
	type fields struct {
		api SSMGetParamAPI
	}
	type args struct {
		ctx            context.Context
		name           string
		withDecryption bool
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       *ssm.GetParameterOutput
		wantErrMsg string
		setup      func(f *fields)
	}{
		{
			name: "ShouldErrorOnGetParameter",
			args: args{
				ctx:            context.Background(),
				name:           "foo_param",
				withDecryption: true,
			},
			wantErrMsg: "foo error",
			setup: func(f *fields) {
				apiMock := mocks.NewSSMGetParamAPI(t)

				apiMock.
					EXPECT().
					GetParameter(context.Background(), paramInput).
					Return(nil, fmt.Errorf("foo error")).
					Once()

				f.api = apiMock
			},
		},
		{
			name: "ShouldGetParameter",
			args: args{
				ctx:            context.Background(),
				name:           "foo_param",
				withDecryption: true,
			},
			want: paramOutput,
			setup: func(f *fields) {
				apiMock := mocks.NewSSMGetParamAPI(t)

				apiMock.
					EXPECT().
					GetParameter(context.Background(), paramInput).
					Return(paramOutput, nil).
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
			got, err := c.GetParameter(tt.args.ctx, tt.args.name, tt.args.withDecryption)
			if err != nil {
				assert.EqualError(t, err, tt.wantErrMsg)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
