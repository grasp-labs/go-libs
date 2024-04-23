docker-up:
	@echo "Building and starting local environment"
	@docker-compose -f ./docker/docker-compose.yaml up --build -d

test:
	@echo "Running Golang unit tests"
	@go test -v -short ./... -skip Example

test-integration:
	@echo "Running Golang integration tests"
	@go test -run Integration -v ./...

lint:
	@echo "Running Golang Linter"
	@golangci-lint run
	@echo "Running Markdown Linter"
	@markdownlint *.md
