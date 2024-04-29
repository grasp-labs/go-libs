#!/bin/bash

echo "Initializing Localstack infra!"

awslocal configure set aws_access_key_id 'dummy' --profile test-profile
awslocal configure set aws_secret_access_key 'dummy' --profile test-profile
awslocal configure set region 'eu-central-1' --profile test-profile
awslocal configure set output 'table' --profile test-profile
awslocal --endpoint-url=http://localhost:4566 ssm put-parameter --name "AUTH_JWT_PUBLIC_KEY_TEST" --profile test-profile --type SecureString --value "1234" --overwrite --region "eu-north-1" --output table
awslocal --endpoint-url=http://localhost:4566 sqs create-queue --queue-name daas-service-cost-handler-usage-queue-test --profile test-profile --region "eu-north-1" --output table
awslocal --endpoint-url=http://localhost:4566 dynamodb create-table \
    --table-name audit_test \
    --key-schema AttributeName=id,KeyType=HASH \
    --attribute-definitions AttributeName=id,AttributeType=S \
    --billing-mode PAY_PER_REQUEST \
    --region "eu-north-1" \
    --profile test-profile \
    --output table

awslocal --endpoint-url=http://localhost:4566 dynamodb create-table \
    --table-name workflow_status_test \
    --key-schema AttributeName=workflow_id,KeyType=HASH \
    --attribute-definitions AttributeName=workflow_id,AttributeType=S \
    --billing-mode PAY_PER_REQUEST \
    --region "eu-north-1" \
    --profile test-profile \
    --output table

awslocal dynamodb put-item \
    --table-name workflow_status_test \
    --item '{"workflow_id":{"S":"d3620960-d2cc-4419-930d-78a56b92a206"}, "status":{"S":"in progress"}}' \
    --region "eu-north-1" \
     --profile test-profile \
    --output table

#awslocal dynamodb describe-table \
#    --table-name workflow_status_test \
#    --region "eu-north-1"

#awslocal dynamodb get-item \
#    --table-name workflow_status_test \
#    --key '{"id":{"S":"1"}}' \
#    --region "eu-north-1"

awslocal --endpoint-url=http://localhost:4566 s3api create-bucket \
    --bucket test \
    --profile test-profile \
    --create-bucket-configuration LocationConstraint=eu-central-1 \
    --region eu-central-1 \
    --output table

awslocal --endpoint-url=http://localhost:4566 s3 \
    cp /tmp/s3/dataset.json s3://test/tenant/c15e32af-71db-4fda-b4e6-2831b1f2b044/workflows/dataset/dataset.json