# Speed Running REST APIs
## An AI-Augmented, go generate-Powered Workflow

---

## About This Talk

**Goal**: Show you how to build production-ready Go REST APIs at unprecedented speed

**Method**: Combine OpenAPI specs + code generation + AI assistance

**Result**: Focus on business logic, not boilerplate

---

## The Problem

Building REST APIs traditionally involves:

- ❌ Writing repetitive CRUD boilerplate
- ❌ Manual type definitions across layers
- ❌ Tedious request/response marshaling
- ❌ Time-consuming test setup
- ❌ Documentation that falls out of sync

**We spend 80% of our time on the boring 20%**

---

## The Solution: A New Paradigm

```
OpenAPI Spec → Code Generation → Business Logic
     ↑              ↑                  ↑
   AI Assist    go generate        Human Focus
```

**Core Principles:**

1. **OpenAPI First, Code Later** - Single source of truth
2. **AI as Force Multiplier** - Automate the tedious
3. **go generate as Orchestrator** - One command to rule them all
4. **Focus on Business Logic** - Your competitive advantage

---

## The Workflow: Three Steps

### 1️⃣ AI-Assisted API Definition
### 2️⃣ Orchestrate with go generate
### 3️⃣ Hyper-Focused Business Logic

---

## Step 1: AI-Assisted API Definition

### The Old Way
```yaml
# Manually writing 200+ lines of YAML...
openapi: 3.0.0
info:
  title: User API
  version: 1.0.0
paths:
  /users/{id}:
    get:
      # ... 50 more lines per endpoint
```

### The New Way
**Prompt AI:** *"Create an openapi.yaml for a RESTful user API..."*

⚡ **Result**: Well-structured, syntactically correct spec in seconds

---

## Step 1: Example AI Prompt

```
Create an openapi.yaml for a RESTful user API. 
It should have:
- GET /users/{id} to retrieve a user
- POST /users to create a new user
- The user object should have id, name, and email
- Ensure name and email are required for POST
```

**AI Output**: Complete OpenAPI 3.0 specification with:
- Proper schemas
- Request/response definitions
- Validation rules
- Error responses

---

## Step 1: Database Schema Generation

**Prompt AI:** *"Based on the User schema from our OpenAPI spec, write the PostgreSQL DDL and sqlc queries"*

```sql
-- AI generates this:
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (name, email) 
VALUES ($1, $2) RETURNING *;
```

---

## Step 2: Orchestrate with go generate

### Single Command, Complete Backend

```go
//go:generate oapi-codegen -config config.yaml openapi.yaml
//go:generate sqlc generate
```

**What happens:**
1. `oapi-codegen` → Go types + server interface
2. `sqlc` → Type-safe database layer

```bash
$ go generate ./...
```

✅ **Done. Your API layer and data layer are ready.**

---

## Step 2: Generated Code Structure

```
.
├── openapi.yaml           # Source of truth
├── gen.go                 # go:generate directives
├── api/
│   └── generated.go       # oapi-codegen output
├── db/
│   ├── queries.sql        # sqlc input
│   ├── sqlc.yaml          # sqlc config
│   └── generated.go       # sqlc output
└── service/
    └── user_service.go    # Your business logic
```

---

## Step 2: What oapi-codegen Generates

```go
// Generated types from OpenAPI spec
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type CreateUserRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

// Generated server interface
type ServerInterface interface {
    GetUser(w http.ResponseWriter, r *http.Request, id int)
    CreateUser(w http.ResponseWriter, r *http.Request)
}
```

---

## Step 2: AI-Generated Scaffolding

**Prompt AI:** *"Generate the Go implementation for our oapi-codegen server interface..."*

```go
// AI generates this boilerplate:
type Server struct {
    userService *service.UserService
}

func (s *Server) GetUser(w http.ResponseWriter, r *http.Request, id int) {
    user, err := s.userService.GetUserByID(r.Context(), id)
    if err != nil {
        // Error handling...
    }
    json.NewEncoder(w).Encode(user)
}

func (s *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
    // Request parsing, validation, service call...
}
```

---

## Step 3: Hyper-Focused Business Logic

### You Only Write This Part

```go
func (s *UserService) CreateUser(ctx context.Context, req CreateUserRequest) (*User, error) {
    // ✅ This is where YOU add value
    
    // Check for duplicate email
    existing, _ := s.db.GetUserByEmail(ctx, req.Email)
    if existing != nil {
        return nil, ErrDuplicateEmail
    }
    
    // Call external service for email verification
    if err := s.emailVerifier.Verify(req.Email); err != nil {
        return nil, ErrInvalidEmail
    }
    
    // Complex business rules
    if s.requiresApproval(req) {
        return s.createPendingUser(ctx, req)
    }
    
    return s.db.CreateUser(ctx, req)
}
```

---

## Step 3: AI-Powered Testing

**Prompt AI:** *"Write comprehensive unit tests for the CreateUser function..."*

```go
// AI generates this:
func TestCreateUser_Success(t *testing.T) {
    // Setup mocks, test cases, assertions...
}

func TestCreateUser_DuplicateEmail(t *testing.T) {
    // Test error handling...
}

func TestCreateUser_InvalidEmail(t *testing.T) {
    // Test validation...
}

func TestCreateUser_RequiresApproval(t *testing.T) {
    // Test business logic branches...
}
```

**Coverage**: 80%+ without writing a single test manually

---

## Step 3: AI-Powered Documentation

**Prompt AI:** *"Generate GoDoc comments for all exported functions..."*

```go
// CreateUser creates a new user in the system after performing
// validation and duplicate checks. If the user requires approval
// based on business rules, it creates a pending user instead.
//
// Parameters:
//   - ctx: Context for cancellation and timeouts
//   - req: CreateUserRequest containing name and email
//
// Returns:
//   - *User: The created user object
//   - error: ErrDuplicateEmail, ErrInvalidEmail, or database errors
func (s *UserService) CreateUser(ctx context.Context, req CreateUserRequest) (*User, error) {
    // ...
}
```

---

## Live Demo Time! 🚀

Let's build a complete REST API from scratch:

1. **AI generates OpenAPI spec** (30 seconds)
2. **AI generates database schema** (30 seconds)
3. **Run go generate** (5 seconds)
4. **AI scaffolds handlers** (1 minute)
5. **Write business logic** (5 minutes)
6. **AI generates tests** (1 minute)

**Total**: ~8 minutes for a production-ready API

---

## The Results: Before vs After

### Traditional Approach
- ⏱️ **Time**: 2-3 days for basic CRUD API
- 📝 **Lines of Code**: ~2000 lines
- 🐛 **Bugs**: Manual marshaling errors, type mismatches
- 📚 **Documentation**: Often outdated or missing

### AI-Augmented Approach
- ⏱️ **Time**: 2-3 hours for the same API
- 📝 **Lines of Code**: ~500 lines (you write ~200)
- 🐛 **Bugs**: Minimal - generated code is consistent
- 📚 **Documentation**: Auto-generated and always in sync

**10x productivity improvement**

---

## Real-World Benefits

### Type Safety Across Layers
```
OpenAPI Schema → Go Types → Database Models
```
Change the spec, regenerate, compiler catches all issues

### Single Source of Truth
- API contract in `openapi.yaml`
- Frontend can generate TypeScript types
- Backend generates Go types
- Documentation auto-generated

### Predictable Build Process
```bash
go generate ./...
go test ./...
go build
```

---

## Tools & Technologies

### Code Generation
- **oapi-codegen** - OpenAPI → Go types & server interface
- **sqlc** - SQL → Type-safe Go database code

### AI Assistants
- GitHub Copilot
- Gemini Code Assist
- Claude / ChatGPT
- Cursor / Windsurf

### Go Tooling
- `go generate` - Built-in orchestration
- Standard library - `net/http`, `encoding/json`

---

## Best Practices

### 1. OpenAPI First
- Define your API contract before writing code
- Use AI to bootstrap, then refine manually
- Version your OpenAPI spec

### 2. Validate Generated Code
- Review AI output before committing
- Run tests after generation
- Use linters (golangci-lint)

### 3. Keep Business Logic Separate
- Don't modify generated code
- Put your logic in service layers
- Use interfaces for testability

---

## Best Practices (Continued)

### 4. Automate Everything
```go
//go:generate oapi-codegen -config config.yaml openapi.yaml
//go:generate sqlc generate
//go:generate go run scripts/gen_mocks.go
```

### 5. Version Control
- Commit generated code (controversial but practical)
- Or use CI/CD to regenerate
- Document generation steps in README

### 6. Iterate Quickly
- Change OpenAPI spec
- Regenerate
- Compiler tells you what broke
- Fix business logic

---

## Common Pitfalls

### ❌ Don't Do This
- Modifying generated code directly
- Skipping the OpenAPI spec
- Over-relying on AI without review
- Ignoring type safety warnings

### ✅ Do This
- Keep generated code read-only
- Treat OpenAPI as source of truth
- Review and test AI output
- Embrace compiler errors as guides

---

## When to Use This Workflow

### ✅ Perfect For:
- REST APIs with standard CRUD operations
- Microservices with clear contracts
- APIs consumed by multiple clients
- Teams that value type safety

### 🤔 Consider Alternatives:
- Real-time/streaming APIs (use gRPC)
- GraphQL APIs (different tooling)
- Very simple APIs (might be overkill)
- Legacy systems (migration complexity)

---

## Extending the Workflow

### Add More Generators
```go
//go:generate oapi-codegen -config config.yaml openapi.yaml
//go:generate sqlc generate
//go:generate mockgen -source=api/generated.go
//go:generate go run scripts/gen_client.go
```

### Generate Client SDKs
- TypeScript/JavaScript for frontend
- Python for data science teams
- Mobile SDKs (Swift, Kotlin)

### Generate Documentation
- Swagger UI
- Redoc
- Postman collections

---

## The Future: Even Faster

### Emerging Patterns
- **AI-to-AI workflows** - AI generates prompts for other AIs
- **Continuous generation** - Watch mode for specs
- **Multi-language codegen** - One spec, many languages

### What's Next
- Better AI understanding of domain models
- Smarter business logic generation
- AI-powered API design suggestions
- Automated performance optimization

---

## Key Takeaways

1. **OpenAPI is your contract** - Everything flows from it
2. **AI accelerates the boring parts** - Specs, boilerplate, tests
3. **go generate orchestrates** - One command, complete backend
4. **You focus on value** - Business logic, not plumbing
5. **Type safety everywhere** - Compiler is your friend

### The Bottom Line
**This isn't about replacing developers. It's about empowering them.**

---

## Resources

### Code & Examples
- GitHub: [your-repo]/speedrun-rest-apis
- Demo project with full implementation
- AI prompts library

### Tools
- oapi-codegen: github.com/deepmap/oapi-codegen
- sqlc: sqlc.dev
- OpenAPI Spec: spec.openapis.org

### Learning
- OpenAPI Tutorial: swagger.io/docs
- Go Code Generation: go.dev/blog/generate

---

## Thank You! 🎉

### Questions?

**Let's discuss:**
- Your current API development workflow
- Challenges with code generation
- AI tools you're using
- How to adopt this in your team

**Connect:**
- Twitter: @yourhandle
- GitHub: @yourhandle
- Email: you@example.com

---

## Bonus: Demo Repository Structure

```
speedrun-rest-api/
├── README.md
├── openapi.yaml              # Source of truth
├── gen.go                    # go:generate directives
├── config.yaml               # oapi-codegen config
├── go.mod
├── api/
│   └── generated.go          # Generated API types
├── db/
│   ├── migrations/
│   ├── queries.sql
│   ├── sqlc.yaml
│   └── generated.go          # Generated DB code
├── service/
│   ├── user_service.go       # Business logic
│   └── user_service_test.go  # Tests
├── server/
│   ├── server.go             # HTTP server setup
│   └── handlers.go           # Handler implementations
└── cmd/
    └── api/
        └── main.go           # Entry point
```

---

## Appendix: Sample AI Prompts

### For OpenAPI Spec
```
Create an OpenAPI 3.0 specification for a user management API with:
- GET /users - list all users with pagination
- GET /users/{id} - get user by ID
- POST /users - create new user
- PUT /users/{id} - update user
- DELETE /users/{id} - delete user

User schema: id (int), name (string), email (string), 
created_at (datetime), updated_at (datetime)

Include proper error responses (400, 404, 500)
```

---

## Appendix: Sample AI Prompts (Continued)

### For Database Schema
```
Based on this OpenAPI User schema, create:
1. PostgreSQL DDL for a users table with proper indexes
2. sqlc queries for all CRUD operations
3. Include created_at and updated_at with automatic timestamps
4. Add a unique constraint on email
5. Include queries for pagination (LIMIT/OFFSET)
```

### For Handler Implementation
```
Implement the ServerInterface from oapi-codegen for user management.
- Use a UserService for business logic
- Handle errors with proper HTTP status codes
- Map between API types and database models
- Add request validation
- Include logging
- Add comments explaining each handler
```

---

## Appendix: Sample AI Prompts (Continued)

### For Tests
```
Generate comprehensive unit tests for the UserService.CreateUser function.
Include test cases for:
- Successful user creation
- Duplicate email error
- Invalid email format
- Database errors
- Context cancellation

Use table-driven tests and mock the database layer.
Aim for 90%+ code coverage.
```

### For Documentation
```
Generate GoDoc comments for all exported types and functions in the
service package. Include:
- Function purpose and behavior
- Parameter descriptions
- Return value descriptions
- Error conditions
- Usage examples where helpful
```
