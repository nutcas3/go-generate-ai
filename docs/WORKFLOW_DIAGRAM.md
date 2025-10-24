# Workflow Diagram

## The Complete Workflow

```
┌─────────────────────────────────────────────────────────────────┐
│                     STEP 1: AI-ASSISTED DEFINITION              │
└─────────────────────────────────────────────────────────────────┘

    ┌──────────────┐
    │   AI Prompt  │  "Create an openapi.yaml for a user API..."
    └──────┬───────┘
           │
           ▼
    ┌──────────────┐
    │ openapi.yaml │  ← Single Source of Truth
    └──────┬───────┘
           │
           ▼
    ┌──────────────┐
    │   AI Prompt  │  "Create PostgreSQL schema and queries..."
    └──────┬───────┘
           │
           ▼
    ┌──────────────┐
    │  schema.sql  │
    │ queries.sql  │
    └──────┬───────┘
           │
           │
┌──────────┴──────────────────────────────────────────────────────┐
│                     STEP 2: CODE GENERATION                      │
└──────────────────────────────────────────────────────────────────┘
           │
           ▼
    ┌──────────────┐
    │ go generate  │  ← Single Command
    └──────┬───────┘
           │
           ├─────────────────┬─────────────────┐
           │                 │                 │
           ▼                 ▼                 ▼
    ┌─────────────┐   ┌─────────────┐   ┌─────────────┐
    │ oapi-codegen│   │    sqlc     │   │   mockgen   │
    └──────┬──────┘   └──────┬──────┘   └──────┬──────┘
           │                 │                 │
           ▼                 ▼                 ▼
    ┌─────────────┐   ┌─────────────┐   ┌─────────────┐
    │ API Types   │   │ DB Queries  │   │   Mocks     │
    │ Server I/F  │   │ Type-safe   │   │ For Testing │
    └─────────────┘   └─────────────┘   └─────────────┘
           │                 │                 │
           └─────────────────┴─────────────────┘
                             │
┌────────────────────────────┴─────────────────────────────────────┐
│                  STEP 3: BUSINESS LOGIC                          │
└──────────────────────────────────────────────────────────────────┘
                             │
                             ▼
                    ┌─────────────────┐
                    │  You Write This │
                    │                 │
                    │ • Validation    │
                    │ • Business Rules│
                    │ • Integrations  │
                    │ • Domain Logic  │
                    └─────────────────┘
                             │
                             ▼
                    ┌─────────────────┐
                    │   AI Prompt     │  "Generate tests..."
                    └────────┬────────┘
                             │
                             ▼
                    ┌─────────────────┐
                    │  Unit Tests     │
                    │  90%+ Coverage  │
                    └─────────────────┘
```

## Layer Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                         HTTP Layer                               │
│  ┌────────────────────────────────────────────────────────┐    │
│  │  server/server.go                                       │    │
│  │  • Request parsing                                      │    │
│  │  • Response formatting                                  │    │
│  │  • HTTP status codes                                    │    │
│  │  • Error handling                                       │    │
│  └────────────────────────────────────────────────────────┘    │
└───────────────────────────┬─────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────────┐
│                      Service Layer                               │
│  ┌────────────────────────────────────────────────────────┐    │
│  │  service/user_service.go                                │    │
│  │  • Business logic ← YOU WRITE THIS                      │    │
│  │  • Validation                                           │    │
│  │  • Orchestration                                        │    │
│  │  • Domain rules                                         │    │
│  └────────────────────────────────────────────────────────┘    │
└───────────────────────────┬─────────────────────────────────────┘
                            │
                            ▼
┌─────────────────────────────────────────────────────────────────┐
│                      Data Layer                                  │
│  ┌────────────────────────────────────────────────────────┐    │
│  │  db/generated.go (by sqlc)                              │    │
│  │  • Type-safe queries                                    │    │
│  │  • Connection pooling                                   │    │
│  │  • Transaction support                                  │    │
│  └────────────────────────────────────────────────────────┘    │
└───────────────────────────┬─────────────────────────────────────┘
                            │
                            ▼
                    ┌───────────────┐
                    │  PostgreSQL   │
                    └───────────────┘
```

## Type Flow

```
┌──────────────┐
│ openapi.yaml │
└──────┬───────┘
       │
       │ oapi-codegen
       │
       ▼
┌──────────────┐     ┌──────────────┐     ┌──────────────┐
│  API Types   │────▶│ Service Layer│────▶│  DB Types    │
│              │     │              │     │              │
│ type User {  │     │ Map between  │     │ type User {  │
│   ID   int   │     │   layers     │     │   ID   int32 │
│   Name string│     │              │     │   Name string│
│   Email...   │     │              │     │   Email...   │
│ }            │     │              │     │ }            │
└──────────────┘     └──────────────┘     └──────────────┘
       │                                          │
       │                                          │
       └──────────────┬───────────────────────────┘
                      │
                      ▼
              ┌──────────────┐
              │  JSON/SQL    │
              │  Consistent  │
              │  Type-safe   │
              └──────────────┘
```

## Development Cycle

```
    ┌─────────────────┐
    │  Change Needed  │
    └────────┬────────┘
             │
             ▼
    ┌─────────────────┐
    │ Update OpenAPI  │ ← Use AI to help
    └────────┬────────┘
             │
             ▼
    ┌─────────────────┐
    │  go generate    │
    └────────┬────────┘
             │
             ▼
    ┌─────────────────┐
    │ Compiler Errors?│
    └────────┬────────┘
             │
         Yes │ No
             │
    ┌────────┴────────┐
    │                 │
    ▼                 ▼
┌─────────┐    ┌─────────────┐
│  Fix    │    │   Tests     │
│ Errors  │    │   Pass?     │
└────┬────┘    └──────┬──────┘
     │                │
     │            Yes │ No
     │                │
     │         ┌──────┴──────┐
     │         │             │
     │         ▼             ▼
     │    ┌────────┐    ┌────────┐
     │    │ Deploy │    │  Fix   │
     │    └────────┘    │ Tests  │
     │                  └───┬────┘
     │                      │
     └──────────────────────┘
```

## Time Comparison

### Traditional Approach (2-3 days)
```
Day 1:
├─ 2h: Design API
├─ 3h: Write types
├─ 2h: Write handlers
└─ 1h: Debug JSON marshaling

Day 2:
├─ 2h: Write SQL queries
├─ 3h: Write database layer
├─ 2h: Debug SQL errors
└─ 1h: Integration

Day 3:
├─ 3h: Write tests
├─ 2h: Write documentation
├─ 2h: Bug fixes
└─ 1h: Final testing
```

### AI-Augmented Approach (2-3 hours)
```
Hour 1:
├─ 10m: AI generates OpenAPI spec
├─ 10m: AI generates SQL schema
├─ 5m:  go generate
├─ 20m: Review generated code
└─ 15m: AI generates handler scaffolding

Hour 2:
├─ 45m: Write business logic
└─ 15m: AI generates tests

Hour 3:
├─ 30m: Run tests, fix issues
├─ 15m: AI generates documentation
└─ 15m: Final testing
```

## Benefits Visualization

```
Traditional:                AI-Augmented:
┌──────────┐               ┌──────────┐
│          │               │          │
│  80%     │               │  20%     │
│          │               │          │
│Boilerplate│               │Boilerplate│
│          │               │          │
│          │               │(Generated)│
│          │               │          │
├──────────┤               ├──────────┤
│          │               │          │
│  20%     │               │  80%     │
│          │               │          │
│ Business │               │ Business │
│  Logic   │               │  Logic   │
│          │               │          │
└──────────┘               └──────────┘

  Your Time                  Your Time
```

## Key Takeaway

```
┌─────────────────────────────────────────────────────────────┐
│                                                             │
│  OpenAPI Spec (1 file)                                      │
│         ↓                                                   │
│  go generate (1 command)                                    │
│         ↓                                                   │
│  Complete API (100s of lines)                               │
│         ↓                                                   │
│  You focus on business logic                                │
│                                                             │
│  Result: 10x faster development                             │
│                                                             │
└─────────────────────────────────────────────────────────────┘
```
