# Changelog

## v1.2.2 - 2024-04-29

### Fixes

* Tests: update mockery.

## v1.2.1 - 2024-04-29

### Fixes

* DynamoDB: change from any to []map[string]any in Query function as return param.

## v1.2.0 - 2024-04-29

### Enhancements

* DynamoDB: Add Query method.


## v1.1.2 - 2024-04-25

### Fixes

* Tests: Move integration tests to new folder - `tests/integration/integration_test.go`

## v1.1.1 - 2024-04-24

* Docs: fix typos in changelog.

## v1.1.0 - 2024-04-24

### Enhancements

* S3: Add Put Object function.
* S3: Add Delete Object function.
* Makefile: change CLI names.
* CHANGELOG.md: Add changelog file with content.
* README.md: Update file for missing AWS services and how to write integration tests.

## v1.0.2 - 2024-04-23

### Enhancements

* S3: Add Get object function.

## v1.0.1 - 2024-04-23

### Chore

* Update go.mod file

## v1.0.0 - 2024-04-23

### Enhancements

* SMS: Add Get parameter function.
* SQS: Add Send message function.
* DynamoDB: Add put item function.
