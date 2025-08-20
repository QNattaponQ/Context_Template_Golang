name: "User Management REST API with PostgreSQL and GORM"
description: |

## Purpose
Implement a REST API endpoint `/users` for user management operations including creating new users and retrieving user data by ID, using PostgreSQL as the database and GORM as the ORM.

## Core Principles
1. **Context is King**: Include ALL necessary documentation, examples, and caveats
2. **Validation Loops**: Provide executable tests/lints the AI can run and fix
3. **Information Dense**: Use keywords and patterns from the codebase
4. **Progressive Success**: Start simple, validate, then enhance
5. **Global rules**: Be sure to follow all rules in CLAUDE.md

---

## Goal
Build a complete REST API with `/users` endpoint that allows:
- POST `/users` to create new users (accepting JSON with name, email, age)
- GET `/users/{id}` to retrieve user data by ID
- Store data in PostgreSQL using GORM as the ORM
- Follow existing Go patterns from the examples directory

## Why
- Provides user management capabilities for the application
- Demonstrates proper Go web development patterns with GORM and Chi router
- Establishes a foundation for future user-related features
- Integrates with existing codebase architecture and testing patterns

## What
A REST API service that handles user CRUD operations with proper error handling, validation, and database persistence.

### Success Criteria
- [ ] POST `/users` successfully creates users and stores them in PostgreSQL
- [ ] GET `/users/{id}` successfully retrieves user data by ID
- [ ] Proper validation of input data (name, email, age)
- [ ] Comprehensive unit tests with table-driven test patterns
- [ ] Integration tests verify database operations
- [ ] Follows existing code patterns from examples directory

## All Needed Context

### Documentation & References (list all context needed to implement the feature)
```yaml
# MUST READ - Include these in your context window
- url: https://gorm.io/docs/
  why: GORM documentation for database operations, models, and migrations
  
- url: https://github.com/go-chi/chi
  why: Chi router documentation for HTTP routing and middleware
  
- file: examples/server.go
  why: Pattern for HTTP server setup with Chi router and middleware
  
- file: examples/repository.go
  why: Repository pattern for database operations and GORM usage
  
- file: examples/test/test_example.go
  why: Table-driven test patterns and testing style
  
- doc: Go standard library
  section: net/http, encoding/json
  critical: Proper HTTP status codes and JSON response handling
```

### Current Codebase tree (run `tree` in the root of the project) to get an overview of the codebase
```bash
Context_Template_Golang/
├── examples/
│   ├── server.go
│   ├── repository.go
│   └── test/
│       └── test_example.go
├── golang_example/
├── PRPs/
│   ├── templates/
│   │   └── prp_base.md
│   └── INITIAL.md
└── README.md
```

### Desired Codebase tree with files to be added and responsibility of file
```bash
Context_Template_Golang/
├── examples/
│   ├── server.go                    # MODIFY: Add user routes
│   ├── repository.go                # MODIFY: Add user repository methods
│   ├── models/                      # CREATE: User model definition
│   │   └── user.go
│   ├── handlers/                    # CREATE: HTTP handlers for user operations
│   │   └── user_handler.go
│   ├── database/                    # CREATE: Database connection and setup
│   │   └── database.go
│   └── test/
│       ├── test_example.go
│       └── test_user.go             # CREATE: User-specific tests
├── migrations/                      # CREATE: Database schema migration
│   └── 001_create_users_table.sql
├── config/                          # CREATE: Configuration management
│   └── config.go
└── main.go                         # CREATE: Main application entry point
```

### Known Gotchas of our codebase & Library Quirks
```go
// CRITICAL: GORM requires proper struct tags for database operations
// Example: `gorm:"column:user_name"` for custom column names
// Example: `gorm:"uniqueIndex"` for unique constraints

// CRITICAL: Chi router requires proper middleware ordering
// Example: CORS middleware should come before route handlers

// CRITICAL: PostgreSQL connection requires proper connection string format
// Example: "host=localhost user=username password=password dbname=dbname sslmode=disable"

// CRITICAL: GORM v2 requires proper error handling for database operations
// Example: Check for gorm.ErrRecordNotFound when querying by ID
```

## Implementation Blueprint

### Data models and structure

Create the core data models to ensure type safety and consistency.
```go
// User model with GORM tags
type User struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Name      string    `json:"name" gorm:"not null"`
    Email     string    `json:"email" gorm:"uniqueIndex;not null"`
    Age       int       `json:"age" gorm:"not null"`
    CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// Request/Response DTOs
type CreateUserRequest struct {
    Name  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required,email"`
    Age   int    `json:"age" validate:"required,min=1,max=150"`
}

type UserResponse struct {
    ID        uint      `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Age       int       `json:"age"`
    CreatedAt time.Time `json:"created_at"`
}
```

### list of tasks to be completed to fullfill the PRP in the order they should be completed

```yaml
Task 1:
CREATE config/config.go:
  - Database connection configuration
  - Environment variable handling
  - Server configuration

Task 2:
CREATE migrations/001_create_users_table.sql:
  - SQL script for creating users table
  - Proper indexes and constraints

Task 3:
CREATE examples/database/database.go:
  - Database connection setup
  - GORM initialization
  - Connection pooling configuration

Task 4:
CREATE examples/models/user.go:
  - User struct with GORM tags
  - Request/Response DTOs
  - Validation tags

Task 5:
MODIFY examples/repository.go:
  - ADD user repository methods
  - FOLLOW existing repository pattern
  - INCLUDE proper error handling

Task 6:
CREATE examples/handlers/user_handler.go:
  - HTTP handlers for POST and GET operations
  - Input validation
  - Error response formatting

Task 7:
MODIFY examples/server.go:
  - ADD user routes
  - INCLUDE user handler
  - FOLLOW existing middleware pattern

Task 8:
CREATE examples/test/test_user.go:
  - Table-driven tests for user operations
  - Database integration tests
  - Mock database for unit tests

Task 9:
CREATE main.go:
  - Application entry point
  - Database migration
  - Server startup
```

### Per task pseudocode as needed added to each task
```go
// Task 1: Config
type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    ServerPort string
}

func LoadConfig() *Config {
    // PATTERN: Use environment variables with defaults
    // GOTCHA: Always validate required environment variables
}

// Task 2: Database Migration
-- Migration SQL script
-- PATTERN: Use SERIAL for auto-incrementing IDs
-- GOTCHA: Add proper indexes for performance

// Task 3: Database Connection
func ConnectDB(config *Config) (*gorm.DB, error) {
    // PATTERN: Use connection pooling
    // GOTCHA: Handle connection errors gracefully
    // CRITICAL: Set proper GORM configuration
}

// Task 4: User Model
type User struct {
    // PATTERN: Use proper GORM tags
    // GOTCHA: Email must be unique
    // CRITICAL: Age validation
}

// Task 5: Repository Methods
func (r *Repository) CreateUser(user *User) error {
    // PATTERN: Use existing repository structure
    // GOTCHA: Handle duplicate email errors
}

func (r *Repository) GetUserByID(id uint) (*User, error) {
    // PATTERN: Return gorm.ErrRecordNotFound if not found
    // GOTCHA: Proper error handling
}

// Task 6: HTTP Handlers
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    // PATTERN: Validate input first
    // PATTERN: Use existing error response format
    // GOTCHA: Set proper HTTP status codes
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    // PATTERN: Extract ID from URL parameters
    // PATTERN: Handle not found cases
}

// Task 7: Server Routes
func (s *Server) setupRoutes() {
    // PATTERN: Use existing Chi router setup
    // PATTERN: Include proper middleware
    // CRITICAL: Route parameter extraction
}

// Task 8: Tests
func TestCreateUser(t *testing.T) {
    // PATTERN: Use table-driven tests
    // PATTERN: Test both success and error cases
    // CRITICAL: Clean up test data
}

// Task 9: Main Application
func main() {
    // PATTERN: Load configuration
    // PATTERN: Initialize database
    // PATTERN: Run migrations
    // PATTERN: Start server
}
```

### Integration Points
```yaml
DATABASE:
  - migration: "Create users table with proper schema"
  - index: "CREATE UNIQUE INDEX idx_users_email ON users(email)"
  - connection: "PostgreSQL with GORM v2"
  
CONFIG:
  - add to: config/config.go
  - pattern: "Environment variable configuration with defaults"
  
ROUTES:
  - add to: examples/server.go
  - pattern: "Chi router with user endpoints"
  - middleware: "CORS, logging, recovery"
```

## Validation Loop

### Level 1: Syntax & Style
```bash
# Run these FIRST - fix any errors before proceeding
go fmt ./...
go vet ./...
golangci-lint run

# Expected: No errors. If errors, READ the error and fix.
```

### Level 2: Unit Tests each new feature/file/function use existing test patterns
```go
// CREATE test_user.go with these test cases:
func TestCreateUser(t *testing.T) {
    tests := []struct {
        name    string
        input   CreateUserRequest
        wantErr bool
    }{
        {
            name: "valid user",
            input: CreateUserRequest{
                Name:  "John Doe",
                Email: "john@example.com",
                Age:   30,
            },
            wantErr: false,
        },
        {
            name: "invalid email",
            input: CreateUserRequest{
                Name:  "John Doe",
                Email: "invalid-email",
                Age:   30,
            },
            wantErr: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}

func TestGetUserByID(t *testing.T) {
    // Test successful retrieval
    // Test user not found
    // Test invalid ID format
}
```

```bash
# Run and iterate until passing:
go test ./examples/test/ -v
# If failing: Read error, understand root cause, fix code, re-run
```

### Level 3: Integration Test
```bash
# Start the service
go run main.go

# Test the endpoints
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com", "age": 30}'

curl http://localhost:8080/users/1

# Expected: Proper JSON responses
# If error: Check logs for stack trace
```

## Final validation Checklist
- [ ] All tests pass: `go test ./... -v`
- [ ] No linting errors: `golangci-lint run`
- [ ] No compilation errors: `go build ./...`
- [ ] Manual test successful: [specific curl commands]
- [ ] Error cases handled gracefully
- [ ] Database operations work correctly
- [ ] Proper HTTP status codes returned
- [ ] Input validation working
- [ ] Documentation updated if needed

## Codebase TREE
.
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── http/
│   │   ├── router.go
│   │   └── server.go
│   ├── db/
│   │   └── connect.go
│   ├── user/
│   │   ├── model.go
│   │   ├── repository.go
│   │   ├── service.go
│   │   └── handler.go
│   └── shared/
│       ├── errs/
│       │   └── errors.go
│       └── responses/
│           └── json.go
├── pkg/                # โค้ดที่อาจ reused ภายนอกได้ (optional)
├── configs/
│   └── config.yaml     # ตั้งค่าพื้นฐาน (port, db dsn)
├── migrations/
│   └── 0001_init.sql   # สคริปต์สร้างตาราง
├── examples/
│   ├── README.md
│   ├── repository.go
│   ├── server.go
│   └── tests/
│       └── test_example.go
├── PRPs/               # โฟลเดอร์ให้ /generate-prp วางไฟล์แผน
├── .claude/
│   └── commands/
│       ├── generate-prp.md
│       └── execute-prp.md
├── .golangci.yml
├── .env.example
├── docker-compose.yml
├── Dockerfile
├── Makefile
├── go.mod
├── go.sum              # (สร้างตอน go get)
├── CLAUDE.md
└── INITIAL.md
---

## Anti-Patterns to Avoid
- ❌ Don't create new patterns when existing ones work
- ❌ Don't skip validation because "it should work"  
- ❌ Don't ignore failing tests - fix them
- ❌ Don't use sync functions in async context
- ❌ Don't hardcode values that should be config
- ❌ Don't catch all exceptions - be specific
- ❌ Don't ignore GORM error handling
- ❌ Don't forget to add proper database indexes
- ❌ Don't skip input validation
- ❌ Don't use wrong HTTP status codes
