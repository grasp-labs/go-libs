package integrationtests

import (
	"context"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	sqstypes "github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/grasp-labs/go-libs/aws/dynamodb"
	"github.com/grasp-labs/go-libs/aws/paramstore"
	"github.com/grasp-labs/go-libs/aws/sqs"
	"github.com/stretchr/testify/assert"
)

const (
	paramName = "AUTH_JWT_PUBLIC_KEY_TEST"
)

func TestSSMIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	if err := os.Setenv("BUILDING_MODE", "test"); err != nil {
		assert.NoError(t, err)
		return
	}

	client, err := paramstore.NewClient(context.Background())
	if err != nil {
		assert.NoError(t, err)
		return
	}

	parameter, err := client.GetParameter(context.Background(), paramName, true)
	if err != nil {
		assert.NoError(t, err)
		return
	}

	var paramOutput = &ssm.GetParameterOutput{
		Parameter: &types.Parameter{
			Value: aws.String("1234"),
		},
		ResultMetadata: middleware.Metadata{},
	}

	assert.Equal(t, paramOutput.Parameter.Value, parameter.Parameter.Value)
}

func TestSQSIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	if err := os.Setenv("BUILDING_MODE", "test"); err != nil {
		assert.NoError(t, err)
		return
	}

	client, err := sqs.NewClient(context.Background(), "daas-service-cost-handler-usage-queue-test")
	if err != nil {
		assert.NoError(t, err)
		return
	}

	input := map[string]sqstypes.MessageAttributeValue{
		"product_id": {
			DataType:    aws.String("String"),
			StringValue: aws.String("foo_bar"),
		},
	}

	if err := client.SendMsg(context.Background(), input); err != nil {
		assert.NoError(t, err)
		return
	}
}

func TestDynamoDB(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	if err := os.Setenv("BUILDING_MODE", "test"); err != nil {
		assert.NoError(t, err)
		return
	}

	client, err := dynamodb.NewClient(context.Background())
	if err != nil {
		assert.NoError(t, err)
		return
	}

	type item struct {
		ID       string `json:"id" dynamodbav:"id"`
		TenantID string `json:"tenant_id" dynamodbav:"tenant_id"`
	}

	if err := client.PutItem(context.Background(), "audit_test", item{
		ID:       "foo-id",
		TenantID: "foo-tenant",
	}); err != nil {
		assert.NoError(t, err)
		return
	}
}
