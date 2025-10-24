# Speaker Notes: Speed Running REST APIs

## Talk Duration: 45 minutes
- Introduction: 5 minutes
- Workflow Overview: 10 minutes
- Live Demo: 20 minutes
- Q&A: 10 minutes

---

## Pre-Talk Checklist

### Technical Setup
- [ ] PostgreSQL running locally
- [ ] Database `speedrun_api` created and schema loaded
- [ ] Demo project dependencies installed
- [ ] `oapi-codegen` and `sqlc` installed
- [ ] Terminal with large font size
- [ ] AI assistant configured
- [ ] Postman/curl commands prepared

---

## Key Messages

1. **OpenAPI is your single source of truth** - Everything flows from it
2. **AI accelerates the boring parts** - Specs, boilerplate, tests
3. **go generate orchestrates** - One command, complete backend
4. **You focus on value** - Business logic, not plumbing
5. **10x productivity is real** - Not hyperbole

---

## Demo Script (20 minutes)

### Part 1: Generate OpenAPI Spec (3 min)
- Show AI prompt
- Generate openapi.yaml
- Quick review

### Part 2: Generate Database Schema (2 min)
- Prompt for schema.sql and queries.sql
- Show generated files

### Part 3: Run go generate (2 min)
```bash
go generate ./...
```
- Show generated api/ and db/ code

### Part 4: Implement Handlers (5 min)
- Walk through server/server.go
- Show mapping between layers

### Part 5: Business Logic (3 min)
- Open service/user_service.go
- Highlight CreateUser business rules

### Part 6: Run the API (3 min)
```bash
go run cmd/api/main.go

# In another terminal
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Demo User", "email": "demo@example.com"}'

curl http://localhost:8080/users/1
```

### Part 7: Show Tests (2 min)
```bash
go test -v ./service/
go test -cover ./...
```

---

## Common Questions

**Q: What about GraphQL?**
A: GraphQL has its own codegen. This is for REST, but principles apply.

**Q: How do you handle breaking changes?**
A: OpenAPI versioning + compiler catches issues on regeneration.

**Q: What if AI generates buggy code?**
A: Always review and test. AI is a starting point.

**Q: Can this work with existing projects?**
A: Yes! Generate OpenAPI from existing code, then adopt gradually.

**Q: What about auth?**
A: Define security schemes in OpenAPI, implement as middleware.

**Q: Database migrations?**
A: Use golang-migrate or similar. Keep schema.sql as source of truth.

**Q: Learning curve?**
A: If you know Go and REST, productive in a day.

**Q: Microservices?**
A: Perfect! Each service gets its own OpenAPI spec.

---

## Backup Slides

Have ready in case of technical issues:
- Screenshots of generated code
- Pre-recorded demo video
- Curl command outputs

---

## Post-Talk

- Share GitHub repo link
- Offer to connect on social media
- Collect feedback
