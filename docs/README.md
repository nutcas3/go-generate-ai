# Speed Running REST APIs: An AI-Augmented, go generate-Powered Workflow

A complete technical talk demonstrating how to build production-ready Go REST APIs at 10x speed using OpenAPI specifications, code generation, and AI assistance.

## 📋 Contents

- **SLIDES.md** - Complete slide deck in Markdown format
- **SPEAKER_NOTES.md** - Detailed speaker notes and demo script
- **AI_PROMPTS.md** - Library of AI prompts used in the workflow
- **demo/** - Complete working demo project

## 🎯 Talk Overview

This talk demonstrates a cutting-edge workflow for building robust Go REST APIs by combining:
- OpenAPI specification as single source of truth
- Code generation with oapi-codegen and sqlc
- AI assistance for rapid development
- go generate for orchestration

**Result**: Build production-ready APIs in hours instead of days.

## 🚀 Quick Start

### For Presenters

1. **Review the slides**
   ```bash
   cat SLIDES.md
   ```

2. **Read speaker notes**
   ```bash
   cat SPEAKER_NOTES.md
   ```

3. **Set up the demo**
   ```bash
   cd demo
   make deps          # Install dependencies
   make docker-up     # Start PostgreSQL
   make migrate       # Run migrations
   make generate      # Generate code
   make run          # Start the API
   ```

4. **Test the demo**
   ```bash
   # In another terminal
   curl http://localhost:8080/users
   ```

### For Attendees

1. **Clone the demo project**
   ```bash
   cd demo
   ```

2. **Follow the README**
   ```bash
   cat demo/README.md
   ```

3. **Try the workflow yourself**
   - Modify `openapi.yaml`
   - Run `go generate ./...`
   - See what changes

## 📚 Key Concepts

### 1. OpenAPI First, Code Later
The OpenAPI specification (`openapi.yaml`) is your single source of truth. Everything else is generated from it.

### 2. AI as Force Multiplier
Use AI to generate:
- OpenAPI specifications
- Database schemas
- Boilerplate code
- Tests
- Documentation

### 3. go generate as Orchestrator
One command generates everything:
```go
//go:generate oapi-codegen -config config.yaml openapi.yaml
//go:generate sqlc generate
```

### 4. Focus on Business Logic
Spend time on what matters:
- ✅ Business rules
- ✅ Complex validation
- ✅ External integrations
- ❌ JSON marshaling
- ❌ HTTP routing
- ❌ SQL queries

## 🛠️ Tools Used

- **[oapi-codegen](https://github.com/deepmap/oapi-codegen)** - Generate Go types and server interface from OpenAPI
- **[sqlc](https://sqlc.dev)** - Generate type-safe Go code from SQL
- **[chi](https://github.com/go-chi/chi)** - Lightweight HTTP router
- **PostgreSQL** - Database
- **AI Assistants** - GitHub Copilot, Cursor, Claude, etc.

## 📊 Results

### Before (Traditional Approach)
- ⏱️ **Time**: 2-3 days for basic CRUD API
- 📝 **Lines of Code**: ~2000 lines
- 🐛 **Bugs**: Manual marshaling errors, type mismatches
- 📚 **Documentation**: Often outdated

### After (AI-Augmented Approach)
- ⏱️ **Time**: 2-3 hours for the same API
- 📝 **Lines of Code**: ~500 lines (you write ~200)
- 🐛 **Bugs**: Minimal - generated code is consistent
- 📚 **Documentation**: Auto-generated, always in sync

**10x productivity improvement**

## 🎓 Learning Path

1. **Understand OpenAPI** - Learn the specification format
2. **Try oapi-codegen** - Generate code from a simple spec
3. **Add sqlc** - Generate database layer
4. **Use AI** - Let AI write the specs and queries
5. **Iterate** - Change spec, regenerate, fix errors

## 🔗 Resources

### Documentation
- [OpenAPI Specification](https://spec.openapis.org/oas/v3.0.0)
- [oapi-codegen Documentation](https://github.com/deepmap/oapi-codegen)
- [sqlc Documentation](https://docs.sqlc.dev/)
- [go generate](https://go.dev/blog/generate)

### Tools
- [Swagger Editor](https://editor.swagger.io/) - Edit OpenAPI specs
- [Swagger UI](https://swagger.io/tools/swagger-ui/) - Interactive API docs
- [Postman](https://www.postman.com/) - API testing

### AI Assistants
- [GitHub Copilot](https://github.com/features/copilot)
- [Cursor](https://cursor.sh/)
- [Windsurf](https://codeium.com/windsurf)

## 📁 Project Structure

```
talks/
├── README.md              # This file
├── SLIDES.md             # Presentation slides
├── SPEAKER_NOTES.md      # Speaker notes and demo script
├── AI_PROMPTS.md         # AI prompt library
└── demo/                 # Complete demo project
    ├── README.md
    ├── openapi.yaml      # API specification
    ├── config.yaml       # oapi-codegen config
    ├── gen.go           # go:generate directives
    ├── Makefile         # Build commands
    ├── docker-compose.yml
    ├── Dockerfile
    ├── api/             # Generated API code
    ├── db/              # Database schema and queries
    ├── service/         # Business logic
    ├── server/          # HTTP handlers
    └── cmd/api/         # Entry point
```

## 🎤 Giving This Talk

### Prerequisites
- Go 1.21+
- PostgreSQL 14+
- oapi-codegen installed
- sqlc installed
- AI assistant configured

### Setup (30 minutes before)
1. Test PostgreSQL connection
2. Run through demo once
3. Prepare backup slides/screenshots
4. Test screen sharing
5. Have curl commands ready

### During the Talk
- Start with the problem (boring boilerplate)
- Show the solution (AI + codegen)
- Live demo (20 minutes)
- Emphasize: AI assists, doesn't replace
- Take questions

### After the Talk
- Share GitHub repo
- Offer to help attendees
- Collect feedback

## 🤝 Contributing

Found an issue or have an improvement? Feel free to:
- Open an issue
- Submit a pull request
- Share your experience using this workflow

## 📄 License

MIT License - feel free to use this talk and demo for your own presentations.

## 🙏 Acknowledgments

- OpenAPI Initiative for the specification
- oapi-codegen and sqlc maintainers
- The Go team for go generate
- AI companies for making this workflow possible

## 📧 Contact

Questions about the talk or workflow? Reach out:
- GitHub: [your-github]
- Twitter: [@your-handle]
- Email: you@example.com

---

**Remember**: This workflow isn't about replacing developers. It's about empowering them to focus on what matters: solving real problems and building great products.
