# AI Prompts Library

This document contains the AI prompts used to generate various parts of the demo project.

## OpenAPI Specification

### Basic User API
```
Create an openapi.yaml for a RESTful user API. It should have:
- GET /users - list all users with pagination (limit and offset query params)
- GET /users/{id} - get user by ID
- POST /users - create new user
- PUT /users/{id} - update user
- DELETE /users/{id} - delete user

The user object should have:
- id (integer)
- name (string, required)
- email (string, email format, required)
- created_at (datetime)
- updated_at (datetime)

Include proper error responses (400, 404, 409, 500) with error schema.
Use OpenAPI 3.0 format.
```

### Adding More Endpoints
```
Add the following endpoints to the existing openapi.yaml:
- GET /users/search?q={query} - search users by name or email
- POST /users/{id}/activate - activate a user account
- POST /users/{id}/deactivate - deactivate a user account

Include proper request/response schemas and error handling.
```

## Database Schema

### PostgreSQL Schema
```
Based on the User schema from our OpenAPI spec, create:
1. PostgreSQL DDL for a users table with proper data types
2. Include indexes on email (unique) and id
3. Add created_at and updated_at with automatic timestamps
4. Include comments explaining each column
```

### sqlc Queries
```
Create sqlc queries for the users table:
1. GetUserByID - get single user by ID
2. GetUserByEmail - get single user by email
3. ListUsers - get paginated list of users (LIMIT/OFFSET)
4. CountUsers - get total count of users
5. CreateUser - insert new user and return the created record
6. UpdateUser - update user by ID and return updated record
7. DeleteUser - delete user by ID

Use PostgreSQL syntax and follow sqlc conventions.
```

## Service Layer

### Business Logic Implementation
```
Generate a Go service layer for user management with the following:
- UserService struct that wraps database queries
- GetUserByID(ctx, id) method
- ListUsers(ctx, limit, offset) method with pagination
- CreateUser(ctx, name, email) method with:
  * Input validation
  * Duplicate email check
  * Business rule: special handling for corporate emails (@company.com)
- UpdateUser(ctx, id, name, email) method
- DeleteUser(ctx, id) method

Include:
- Proper error handling with custom error types
- GoDoc comments for all exported functions
- Context support for cancellation
- Logging points for important operations
```

### Error Types
```
Create custom error types for the user service:
- ErrUserNotFound - when user doesn't exist
- ErrDuplicateEmail - when email already exists
- ErrInvalidInput - when input validation fails

Use Go's errors package conventions.
```

## HTTP Handlers

### Server Implementation
```
Implement the ServerInterface from oapi-codegen for user management.
The server should:
- Have a UserService field
- Implement all handler methods (GetUser, ListUsers, CreateUser, UpdateUser, DeleteUser)
- Map between API types (from oapi-codegen) and database types (from sqlc)
- Handle errors with proper HTTP status codes:
  * 200 for successful GET/PUT
  * 201 for successful POST
  * 204 for successful DELETE
  * 400 for invalid input
  * 404 for not found
  * 409 for conflicts
  * 500 for server errors
- Include request/response logging
- Add GoDoc comments explaining each handler

Use encoding/json for marshaling and net/http for responses.
```

### Router Setup
```
Create a SetupRouter function that:
- Uses chi router
- Adds middleware for logging, recovery, and request ID
- Registers the oapi-codegen handlers
- Returns an http.Handler

Include proper middleware ordering.
```

## Testing

### Unit Tests
```
Generate comprehensive unit tests for the UserService.CreateUser function.
Include test cases for:
- Successful user creation
- Duplicate email error
- Invalid input (empty name, empty email)
- Database errors
- Context cancellation

Use:
- Table-driven tests where appropriate
- Mock database layer
- testify/assert for assertions
- Descriptive test names
- Setup and teardown functions

Aim for 90%+ code coverage.
```

### Mock Database
```
Create a mock implementation of the database Queries interface for testing.
Include:
- MockQueries struct with function fields for each method
- Default implementations that return empty/error values
- Helper methods to set up test scenarios
- Clear documentation on how to use in tests
```

### Integration Tests
```
Create integration tests for the HTTP handlers that:
- Use httptest to create test server
- Test all endpoints (GET, POST, PUT, DELETE)
- Verify status codes and response bodies
- Test error cases (404, 400, 409)
- Use a mock database layer
- Include table-driven tests for multiple scenarios
```

## Documentation

### GoDoc Comments
```
Generate GoDoc comments for all exported types and functions in the service package.
For each function, include:
- Brief description of what it does
- Parameter descriptions
- Return value descriptions
- Error conditions and what errors can be returned
- Usage examples where helpful

Follow Go documentation conventions.
```

### README
```
Create a comprehensive README.md for the project that includes:
- Project overview and goals
- Prerequisites and installation instructions
- Quick start guide
- API endpoint documentation with curl examples
- Development workflow (how to modify the API)
- Testing instructions
- Deployment considerations
- Troubleshooting section
- Links to relevant documentation

Use clear formatting with code blocks and examples.
```

## Advanced Prompts

### Adding Authentication
```
Add JWT authentication to the OpenAPI spec:
1. Define a security scheme for Bearer token
2. Add security requirements to protected endpoints
3. Create a POST /auth/login endpoint that returns a JWT
4. Create a POST /auth/register endpoint

Then generate:
- Go middleware for JWT validation
- Auth service with login/register methods
- Tests for authentication flows
```

### Adding Pagination Metadata
```
Enhance the ListUsers endpoint to return pagination metadata:
- Current page
- Total pages
- Total items
- Has next page
- Has previous page
- Links to next/previous pages (HATEOAS style)

Update the OpenAPI spec and regenerate the code.
```

### Adding Filtering
```
Add filtering capabilities to GET /users:
- Filter by email domain
- Filter by creation date range
- Filter by name (partial match)
- Combine multiple filters

Update OpenAPI spec with query parameters and generate the implementation.
```

## Tips for Better AI Prompts

### Be Specific
❌ "Create a user API"
✅ "Create an OpenAPI 3.0 spec for a user API with GET, POST, PUT, DELETE endpoints"

### Provide Context
❌ "Generate tests"
✅ "Generate unit tests for UserService.CreateUser using table-driven tests and mocks"

### Specify Format
❌ "Create database schema"
✅ "Create PostgreSQL DDL with indexes and constraints, formatted for sqlc"

### Include Requirements
❌ "Add error handling"
✅ "Add error handling with custom error types, proper HTTP status codes, and logging"

### Request Documentation
Always add: "Include comments explaining the code"

### Iterate
Start with a basic prompt, then refine:
1. "Create basic OpenAPI spec"
2. "Add validation rules to the spec"
3. "Add error responses"
4. "Add examples to the spec"
