# Quick Reference Guide

## Essential Commands

### Setup
```bash
# Install tools
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Start database
docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=postgres postgres:15

# Create database
createdb speedrun_api
```

### Development Workflow
```bash
# 1. Edit openapi.yaml (use AI to help)
# 2. Edit db/queries.sql (use AI to help)

# 3. Generate all code
go generate ./...

# 4. Run tests
go test ./...

# 5. Run server
go run cmd/api/main.go
```

### Testing Endpoints
```bash
# Create user
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Test User", "email": "test@example.com"}'

# Get user
curl http://localhost:8080/users/1

# List users
curl http://localhost:8080/users?limit=10&offset=0

# Update user
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Updated Name", "email": "new@example.com"}'

# Delete user
curl -X DELETE http://localhost:8080/users/1
```

## Key AI Prompts

### OpenAPI Spec
```
Create an openapi.yaml for a RESTful [resource] API with:
- GET /[resource] - list with pagination
- GET /[resource]/{id} - get by ID
- POST /[resource] - create
- PUT /[resource]/{id} - update
- DELETE /[resource]/{id} - delete

Include proper schemas, validation, and error responses.
```

### Database Schema
```
Based on the [Resource] schema from our OpenAPI spec, create:
1. PostgreSQL DDL with indexes and constraints
2. sqlc queries for all CRUD operations
```

### Service Layer
```
Generate a Go service layer for [resource] management with:
- [Resource]Service struct
- CRUD methods with proper error handling
- Business logic for [specific rules]
- GoDoc comments
```

### Tests
```
Generate comprehensive unit tests for [function].
Include success and error cases.
Use table-driven tests and mocks.
Aim for 90%+ coverage.
```

## File Structure

```
project/
├── openapi.yaml          # Source of truth
├── config.yaml           # oapi-codegen config
├── gen.go               # go:generate directives
├── go.mod
├── api/
│   └── generated.go     # Generated (don't edit)
├── db/
│   ├── schema.sql
│   ├── queries.sql
│   ├── sqlc.yaml
│   └── generated.go     # Generated (don't edit)
├── service/
│   ├── *_service.go     # Your business logic
│   └── *_service_test.go
├── server/
│   └── server.go        # HTTP handlers
└── cmd/api/
    └── main.go          # Entry point
```

## Common Issues

### "command not found: oapi-codegen"
```bash
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
export PATH=$PATH:$(go env GOPATH)/bin
```

### "cannot connect to database"
```bash
# Check PostgreSQL is running
pg_isready

# Check connection string
export DATABASE_URL="postgres://user:pass@localhost:5432/dbname?sslmode=disable"
```

### "generated code has errors"
```bash
# Regenerate everything
rm -rf api/generated.go db/generated.go
go generate ./...
go mod tidy
```

## Best Practices

1. **Never edit generated code** - It will be overwritten
2. **Always review AI output** - AI makes mistakes
3. **Test after generation** - Run `go test ./...`
4. **Version your OpenAPI spec** - It's your contract
5. **Use meaningful names** - AI works better with clear names
6. **Document your prompts** - Save them for reuse
7. **Commit generated code** - Or regenerate in CI/CD

## Resources

- OpenAPI: https://spec.openapis.org
- oapi-codegen: https://github.com/deepmap/oapi-codegen
- sqlc: https://sqlc.dev
- Swagger Editor: https://editor.swagger.io
