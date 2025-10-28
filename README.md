# Speed Running REST APIs - Demo Project

This is a complete demonstration of the AI-augmented, `go generate`-powered workflow for building REST APIs in Go.

## Overview

This project showcases how to build a production-ready REST API using:
- **OpenAPI 3.0** specification as the single source of truth
- **oapi-codegen** for generating Go types and server interfaces
- **sqlc** for type-safe database access
- **go generate** to orchestrate all code generation
- **AI assistance** to accelerate development

## Project Structure

```
.
├── README.md                  # This file
├── openapi.yaml              # OpenAPI specification (source of truth)
├── config.yaml               # oapi-codegen configuration
├── gen.go                    # go:generate directives
├── go.mod                    # Go module definition
├── api/
│   └── generated.go          # Generated API types and interfaces (by oapi-codegen)
├── db/
│   ├── schema.sql           # Database schema
│   ├── queries.sql          # SQL queries for sqlc
│   ├── sqlc.yaml            # sqlc configuration
│   └── generated.go         # Generated database code (by sqlc)
├── service/
│   ├── user_service.go      # Business logic layer
│   └── user_service_test.go # Unit tests
├── server/
│   └── server.go            # HTTP handlers and routing
└── cmd/
    └── api/
        └── main.go          # Application entry point
```

## Prerequisites

- Go 1.21 or later
- PostgreSQL 14 or later
- [oapi-codegen](https://github.com/deepmap/oapi-codegen): `go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest`
- [sqlc](https://sqlc.dev): `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`

## Quick Start

### 1. Set up the database

```bash
# Create database
createdb speedrun_api

# Run migrations
psql speedrun_api < db/schema.sql
```

### 2. Generate code

```bash
# This single command generates all the boilerplate!
go generate ./...
```

This will:
- Generate Go types and server interface from `openapi.yaml`
- Generate type-safe database code from SQL queries

### 3. Install dependencies

```bash
go mod download
```

### 4. Run the server

```bash
# Set database URL (optional, defaults to localhost)
export DATABASE_URL="postgres://postgres:postgres@localhost:5432/speedrun_api?sslmode=disable"

# Run the server
go run cmd/api/main.go
```

The API will be available at `http://localhost:8080`

## API Endpoints

### List Users
```bash
curl http://localhost:8080/users?limit=10&offset=0
```

### Get User by ID
```bash
curl http://localhost:8080/users/1
```

### Create User
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com"}'
```

### Update User
```bash
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Jane Doe", "email": "jane@example.com"}'
```

### Delete User
```bash
curl -X DELETE http://localhost:8080/users/1
```

## Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

## The Workflow: Step by Step

### Step 1: Define API with OpenAPI (AI-Assisted)

Instead of manually writing the OpenAPI spec, we prompted an AI:

**Prompt:**
```
Create an openapi.yaml for a RESTful user API. It should have:
- GET /users - list users with pagination
- GET /users/{id} - get user by ID
- POST /users - create new user
- PUT /users/{id} - update user
- DELETE /users/{id} - delete user

User schema: id, name, email, created_at, updated_at
Include proper error responses.
```

**Result:** Complete `openapi.yaml` in seconds!

### Step 2: Generate Database Schema (AI-Assisted)

We prompted AI to create the database schema:

**Prompt:**
```
Based on the User schema from our OpenAPI spec, create:
1. PostgreSQL DDL for a users table
2. sqlc queries for all CRUD operations
3. Include proper indexes and constraints
```

**Result:** `db/schema.sql` and `db/queries.sql` ready to use!

### Step 3: Run Code Generation

```bash
go generate ./...
```

This single command:
1. Generates Go types from OpenAPI spec
2. Generates server interface with all handler signatures
3. Generates type-safe database layer from SQL

### Step 4: Implement Business Logic

This is where YOU add value! The generated code handles:
- ✅ Request/response marshaling
- ✅ Type definitions
- ✅ Database queries
- ✅ Basic routing

You focus on:
- ✨ Business rules (e.g., duplicate email checks)
- ✨ Complex validation
- ✨ External service integration
- ✨ Domain-specific logic

See `service/user_service.go` for examples.

### Step 5: AI-Generated Tests

We prompted AI to generate comprehensive tests:

**Prompt:**
```
Generate comprehensive unit tests for the UserService.CreateUser function.
Include test cases for success, duplicate email, invalid input, and database errors.
Use table-driven tests and mocks.
```

**Result:** 90%+ test coverage without writing tests manually!

## Key Benefits

### 🚀 Speed
- From idea to working API in minutes, not days
- AI generates boilerplate instantly
- `go generate` rebuilds everything in seconds

### 🔒 Type Safety
- OpenAPI schema → Go types → Database models
- Compiler catches mismatches immediately
- No runtime surprises

### 📚 Documentation
- OpenAPI spec serves as documentation
- Always in sync with implementation
- Can generate Swagger UI, client SDKs, etc.

### 🧪 Testability
- Clean separation of concerns
- Easy to mock database layer
- AI can generate comprehensive tests

### 🔄 Maintainability
- Single source of truth (OpenAPI spec)
- Change spec, regenerate, fix compilation errors
- Predictable build process

## Modifying the API

Want to add a new field or endpoint? Here's the workflow:

1. **Update `openapi.yaml`** (use AI to help)
2. **Run `go generate ./...`**
3. **Fix compilation errors** (compiler tells you what changed)
4. **Update business logic** as needed
5. **Regenerate tests** with AI

## AI Prompts Library

### For OpenAPI Specs
```
Create an OpenAPI 3.0 specification for [describe your API].
Include [list endpoints and models].
Ensure proper validation rules and error responses.
```

### For Database Schema
```
Based on [schema name] from our OpenAPI spec, create:
1. PostgreSQL DDL with proper indexes
2. sqlc queries for [operations]
3. Include [specific requirements]
```

### For Handler Implementation
```
Implement the ServerInterface from oapi-codegen for [feature].
- Use [ServiceName] for business logic
- Handle errors with proper HTTP status codes
- Map between API types and database models
- Add logging and comments
```

### For Tests
```
Generate comprehensive unit tests for [function name].
Include test cases for:
- [success case]
- [error case 1]
- [error case 2]
Use table-driven tests and mocks. Aim for 90%+ coverage.
```

### For Documentation
```
Generate GoDoc comments for all exported types and functions in [package].
Include function purpose, parameters, return values, and error conditions.
```

## Production Considerations

### Environment Variables
- `DATABASE_URL`: PostgreSQL connection string
- `PORT`: Server port (default: 8080)
- `LOG_LEVEL`: Logging level

### Database Migrations
Consider using a migration tool like:
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [goose](https://github.com/pressly/goose)
- [atlas](https://atlasgo.io)

### Observability
Add:
- Structured logging (e.g., `zerolog`, `zap`)
- Metrics (e.g., Prometheus)
- Tracing (e.g., OpenTelemetry)
- Health checks

### Security
- Input validation (already in OpenAPI spec)
- Authentication/authorization middleware
- Rate limiting
- CORS configuration
- SQL injection protection (sqlc handles this)

### Performance
- Connection pooling (already configured)
- Database indexes (in schema.sql)
- Caching layer (Redis, etc.)
- Request timeouts

## Troubleshooting

### Code generation fails
```bash
# Ensure tools are installed
which oapi-codegen
which sqlc

# Reinstall if needed
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

### Database connection fails
```bash
# Check PostgreSQL is running
pg_isready

# Verify connection string
psql $DATABASE_URL
```

### Tests fail
```bash
# Run with verbose output
go test -v ./...

# Check for missing mocks or incorrect test data
```

## Resources

### Documentation
- [OpenAPI Specification](https://spec.openapis.org/oas/v3.0.0)
- [oapi-codegen](https://github.com/deepmap/oapi-codegen)
- [sqlc](https://docs.sqlc.dev/)
- [go generate](https://go.dev/blog/generate)

### Tools
- [Swagger Editor](https://editor.swagger.io/) - Edit OpenAPI specs
- [Swagger UI](https://swagger.io/tools/swagger-ui/) - Interactive API docs
- [Postman](https://www.postman.com/) - API testing

### AI Assistants
- [GitHub Copilot](https://github.com/features/copilot)
- [Cursor](https://cursor.sh/)
- [Windsurf](https://codeium.com/windsurf)

## License

MIT

## Contributing

This is a demo project for the "Speed Running REST APIs" talk. Feel free to use it as a template for your own projects!

## Questions?

Open an issue or reach out to the talk presenter.
