.PHONY: help generate build run test clean docker-up docker-down migrate

help: 
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

generate: ## Generate code from OpenAPI spec and SQL queries
	go generate ./...

build: generate ## Build the application
	go build -o bin/api cmd/api/main.go

run: ## Run the application
	go run cmd/api/main.go

test: ## Run tests
	go test -v ./...

test-coverage: ## Run tests with coverage
	go test -cover ./...
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

lint: ## Run linter
	golangci-lint run

clean: ## Clean build artifacts
	rm -rf bin/
	rm -f coverage.out coverage.html

docker-up: ## Start PostgreSQL in Docker
	docker run --name speedrun-postgres \
		-e POSTGRES_PASSWORD=postgres \
		-e POSTGRES_DB=speedrun_api \
		-p 5432:5432 \
		-d postgres:15

docker-down: ## Stop PostgreSQL Docker container
	docker stop speedrun-postgres
	docker rm speedrun-postgres

migrate: ## Run database migrations
	psql $$DATABASE_URL -f db/schema.sql

deps: ## Install dependencies
	go mod download
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

validate-openapi: ## Validate OpenAPI specification
	@command -v openapi-generator-cli >/dev/null 2>&1 || { echo "openapi-generator-cli not found. Install it first."; exit 1; }
	openapi-generator-cli validate -i openapi.yaml

dev: docker-up migrate run ## Start development environment
