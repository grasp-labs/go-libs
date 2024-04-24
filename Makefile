.PHONY: docker
docker:
	@echo "Building and starting local environment"
	@docker-compose -f ./docker/docker-compose.yaml up --build -d

.PHONY: test
test:
	@echo "Running Golang unit tests"
	@go test -v -short ./... -skip Example

.PHONY: integration
integration:
	@echo "Running Golang integration tests"
	@go test -run Integration -v ./...

.PHONY: lint
lint:
	@echo "Running Golang Linter"
	@golangci-lint run
	@echo "Running Markdown Linter"
	@markdownlint *.md
