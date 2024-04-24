# AWS libs

In this package, you will find all sharable aws tools.

## Available middlewares

|             | Parameter store | DynamoDB | S3 | SQS |
|-------------|-----------------|----------|----|-----|
| Implemented | ✅               | ✅        | ✅  | ✅   |

## Running aws tools locally

Use Makefile to start docker container with fake AWS environment (Localstack).

## Integration tests

Please write integration tests.

Integration test should:

* Include **Integration** word!
* Include `testing.Short()`!

Example:

```go
package integrationtests

import "testing"


func TestSQSIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	return
}
```
