package paramstore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/ssm"

	"github.com/grasp-labs/go-libs/config"
)

type SSMGetParamAPI interface {
	GetParameter(c context.Context, params *ssm.GetParameterInput, optFns ...func(*ssm.Options)) (*ssm.GetParameterOutput, error)
}

type SSMClient interface {
	GetParameter(c context.Context, name string, withDecryption bool) (*ssm.GetParameterOutput, error)
}

type Client struct {
	api SSMGetParamAPI
}

func NewClientWithAPI(api SSMGetParamAPI) (*Client, error) {
	return &Client{api: api}, nil
}

func NewClient(c context.Context) (*Client, error) {
	cfg, err := config.NewConfig(c)
	if err != nil {
		return nil, err
	}

	return &Client{api: ssm.NewFromConfig(cfg)}, nil
}
