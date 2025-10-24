# Getting Started with Speed Running REST APIs

This guide will help you adopt the AI-augmented, go generate-powered workflow for your own projects.

## Prerequisites

### Required Knowledge
- Basic Go programming
- REST API concepts
- SQL fundamentals
- Command line basics

### Required Tools
- Go 1.21 or later
- PostgreSQL 14 or later
- Git
- A code editor (VS Code, GoLand, etc.)
- An AI assistant (GitHub Copilot, Cursor, Claude, etc.)

## Installation

### 1. Install Go Tools

```bash
# Install oapi-codegen
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

# Install sqlc
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Verify installations
oapi-codegen --version
sqlc version
```

### 2. Set Up PostgreSQL

```bash
# Option 1: Docker (recommended for development)
docker run -d \
  --name dev-postgres \
  -e POSTGRES_PASSWORD=postgres \
  -p 5432:5432 \
  postgres:15

# Option 2: Install locally (macOS)
brew install postgresql@15
brew services start postgresql@15

# Create a test database
createdb my_api_db
```

### 3. Configure Your Editor

#### VS Code
```bash
# Install Go extension
code --install-extension golang.go

# Install AI assistant (choose one)
code --install-extension GitHub.copilot
# or use Cursor editor
```

#### GoLand
- Install GitHub Copilot plugin from marketplace
- Enable Go modules support

## Your First API in 30 Minutes

### Step 1: Create Project Structure (2 minutes)

```bash
# Create project
mkdir my-api && cd my-api
go mod init github.com/yourusername/my-api

# Create directories
mkdir -p api db service server cmd/api
```

### Step 2: Generate OpenAPI Spec with AI (5 minutes)

**Prompt your AI assistant:**
```
Create an openapi.yaml for a RESTful task management API with:
- GET /tasks - list tasks with pagination
- GET /tasks/{id} - get task by ID
- POST /tasks - create task
- PUT /tasks/{id} - update task
- DELETE /tasks/{id} - delete task

Task schema:
- id (integer)
- title (string, required)
- description (string)
- completed (boolean, default false)
- created_at (datetime)
- updated_at (datetime)

Include proper validation and error responses.
```

Save the output as `openapi.yaml`

### Step 3: Generate Database Schema with AI (3 minutes)

**Prompt your AI assistant:**
```
Based on the Task schema from our OpenAPI spec, create:
1. PostgreSQL DDL for a tasks table
2. sqlc queries for all CRUD operations
3. Include indexes and constraints
```

Save DDL as `db/schema.sql` and queries as `db/queries.sql`

### Step 4: Configure Code Generation (2 minutes)

Create `config.yaml`:
```yaml
package: api
generate:
  chi-server: true
  models: true
  embedded-spec: true
output: api/generated.go
```

Create `db/sqlc.yaml`:
```yaml
version: "2"
sql:
  - engine: "postgresql"
    queries: "db/queries.sql"
    schema: "db/schema.sql"
    gen:
      go:
        package: "db"
        out: "db"
        sql_package: "pgx/v5"
        emit_json_tags: true
```

Create `gen.go`:
```go
package main

//go:generate oapi-codegen -config config.yaml openapi.yaml
//go:generate sqlc generate
```

### Step 5: Generate Code (1 minute)

```bash
go generate ./...
go mod tidy
```

You now have:
- `api/generated.go` - API types and server interface
- `db/generated.go` - Type-safe database queries

### Step 6: Implement Service Layer with AI (5 minutes)

**Prompt your AI assistant:**
```
Generate a Go service layer for task management with:
- TaskService struct
- GetTaskByID method
- ListTasks method with pagination
- CreateTask method with validation
- UpdateTask method
- DeleteTask method
- Proper error handling
- GoDoc comments
```

Save as `service/task_service.go`

### Step 7: Implement HTTP Handlers with AI (5 minutes)

**Prompt your AI assistant:**
```
Implement the ServerInterface from oapi-codegen for task management.
- Use TaskService for business logic
- Handle errors with proper HTTP status codes
- Map between API types and database types
- Include logging
```

Save as `server/server.go`

### Step 8: Create Main Entry Point (2 minutes)

Create `cmd/api/main.go`:
```go
package main

import (
    "context"
    "log"
    "net/http"
    "os"

    "github.com/yourusername/my-api/db"
    "github.com/yourusername/my-api/server"
    "github.com/jackc/pgx/v5/pgxpool"
)

func main() {
    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        dbURL = "postgres://postgres:postgres@localhost:5432/my_api_db?sslmode=disable"
    }

    pool, err := pgxpool.New(context.Background(), dbURL)
    if err != nil {
        log.Fatal(err)
    }
    defer pool.Close()

    queries := db.New(pool)
    srv := server.NewServer(queries)
    router := server.SetupRouter(srv)

    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
```

### Step 9: Set Up Database (2 minutes)

```bash
# Create database
createdb my_api_db

# Run migrations
psql my_api_db < db/schema.sql
```

### Step 10: Run and Test (3 minutes)

```bash
# Run server
go run cmd/api/main.go

# In another terminal, test it
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "My first task", "description": "Test the API"}'

curl http://localhost:8080/tasks
```

🎉 **Congratulations!** You just built a complete REST API in 30 minutes!

## Next Steps

### Add Tests

**Prompt your AI assistant:**
```
Generate comprehensive unit tests for TaskService.
Include success and error cases.
Use table-driven tests and mocks.
```

Run tests:
```bash
go test ./...
```

### Add More Features

1. **Authentication**
   - Prompt AI to add JWT auth to OpenAPI spec
   - Generate auth middleware
   - Regenerate code

2. **Filtering**
   - Add query parameters to OpenAPI spec
   - Update queries in `db/queries.sql`
   - Regenerate

3. **Relationships**
   - Add related resources (e.g., users, projects)
   - Update OpenAPI spec
   - Generate code

### Deploy

```bash
# Build
go build -o api cmd/api/main.go

# Deploy to your platform
# - Heroku
# - AWS Lambda
# - Google Cloud Run
# - Your own server
```

## Common Patterns

### Adding a New Endpoint

1. Update `openapi.yaml` (use AI)
2. Run `go generate ./...`
3. Implement business logic in service layer
4. Run tests
5. Deploy

### Adding a New Field

1. Update schema in `openapi.yaml`
2. Update `db/schema.sql` and queries
3. Run `go generate ./...`
4. Fix compilation errors (compiler guides you)
5. Update tests

### Changing Validation Rules

1. Update OpenAPI spec validation rules
2. Run `go generate ./...`
3. Add business logic validation in service layer
4. Test

## Tips for Success

### 1. Start Small
Don't try to build everything at once. Start with:
- 2-3 endpoints
- 1-2 resources
- Basic CRUD operations

Then iterate and add more.

### 2. Use AI Effectively

**Good prompts:**
- Specific: "Create PostgreSQL DDL for users table with email unique constraint"
- Contextual: "Based on the User schema from our OpenAPI spec..."
- Detailed: "Include proper indexes, constraints, and comments"

**Bad prompts:**
- Vague: "Make a database"
- No context: "Create some queries"
- Too broad: "Build an entire API"

### 3. Review Everything

AI is fast but not perfect. Always:
- Review generated OpenAPI specs
- Check generated SQL queries
- Test generated code
- Verify error handling

### 4. Iterate Quickly

The workflow enables rapid iteration:
1. Change spec
2. Regenerate
3. Fix errors
4. Test
5. Repeat

Don't aim for perfection on first try.

### 5. Document Your Prompts

Keep a `prompts.md` file with prompts that worked well. Reuse and refine them.

## Troubleshooting

### "Generated code doesn't compile"
- Check your OpenAPI spec is valid
- Ensure database types match API types
- Run `go mod tidy`

### "AI generated incorrect code"
- Review and fix manually
- Refine your prompt and try again
- Use the corrected version as example for future prompts

### "Tests are failing"
- Check business logic implementation
- Verify mock setup
- Ensure database schema matches queries

## Resources

### Learning
- [OpenAPI Tutorial](https://swagger.io/docs/specification/about/)
- [oapi-codegen Examples](https://github.com/deepmap/oapi-codegen/tree/master/examples)
- [sqlc Documentation](https://docs.sqlc.dev/)

### Tools
- [Swagger Editor](https://editor.swagger.io/) - Validate OpenAPI specs
- [Postman](https://www.postman.com/) - Test APIs
- [TablePlus](https://tableplus.com/) - Database GUI

### Community
- [oapi-codegen Discussions](https://github.com/deepmap/oapi-codegen/discussions)
- [sqlc Discord](https://discord.gg/EyrZkh9)
- [Gophers Slack](https://gophers.slack.com/)

## What's Next?

Once you're comfortable with the basics:

1. **Explore advanced OpenAPI features**
   - Security schemes
   - Webhooks
   - Callbacks

2. **Add more generators**
   - Mock servers
   - Client SDKs
   - Documentation sites

3. **Optimize your workflow**
   - CI/CD integration
   - Pre-commit hooks
   - Automated testing

4. **Share your experience**
   - Write blog posts
   - Give talks
   - Help others adopt the workflow

## Get Help

If you get stuck:
1. Check the demo project in this repo
2. Review the AI prompts library
3. Ask in community forums
4. Open an issue on GitHub

Happy speed running! 🚀
