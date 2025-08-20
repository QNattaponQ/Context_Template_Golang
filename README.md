# User Management REST API

A well-structured Go-based REST API for user management operations using PostgreSQL and GORM, following Go project layout best practices.

## Project Structure

```
.
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internal/
│   ├── http/
│   │   └── server.go            # HTTP server setup and routing
│   ├── db/
│   │   └── connect.go           # Database connection management
│   ├── user/
│   │   ├── model.go             # User data models and DTOs
│   │   ├── repository.go        # Data access layer for users
│   │   ├── handler.go           # HTTP handlers for user operations
│   │   └── user_test.go         # Comprehensive user tests
│   └── shared/
│       ├── errs/
│       │   └── errors.go        # Shared error handling
│       └── responses/
│           └── json.go          # Standardized JSON responses
├── pkg/                         # Reusable packages (optional)
├── configs/
│   └── config.go                # Configuration management
├── migrations/
│   └── 001_create_users_table.sql  # Database schema
├── docker-compose.yml           # PostgreSQL setup
├── go.mod                       # Go module dependencies
└── README.md                    # This file
```

## Features

- **POST** `/users` - Create new users
- **GET** `/users/{id}` - Retrieve user by ID
- PostgreSQL database with GORM ORM
- Comprehensive input validation
- Proper error handling and HTTP status codes
- Table-driven tests with SQLite for testing
- Clean architecture with separation of concerns

## Quick Start

### 1. Start PostgreSQL Database

```bash
# Using Docker Compose (recommended)
docker-compose up -d

# Or manually start PostgreSQL and create database 'userdb'
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Set Environment Variables

```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=password
export DB_NAME=userdb
export SERVER_PORT=8080
```

### 4. Run Database Migration

```bash
psql -h localhost -U postgres -d userdb -f migrations/001_create_users_table.sql
```

### 5. Start the Server

```bash
go run cmd/server/main.go
```

## API Endpoints

### Create User

```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "age": 30
  }'
```

### Get User by ID

```bash
curl http://localhost:8080/users/1
```

### Health Check

```bash
curl http://localhost:8080/health
```

## Testing

Run the comprehensive test suite:

```bash
# Run all tests
go test ./internal/... -v

# Run specific test
go test ./internal/user/ -run TestUserHandler_CreateUser -v

# Run with coverage
go test ./internal/user/ -cover
```

## Architecture

This project follows Go project layout best practices:

- **`cmd/`**: Application entry points
- **`internal/`**: Private application code
- **`pkg/`**: Public libraries that can be imported by other applications
- **`configs/`**: Configuration files
- **`migrations/`**: Database schema changes

## Dependencies

- **Chi Router**: Lightweight HTTP router
- **GORM**: ORM for Go with PostgreSQL support
- **Testify**: Testing utilities
- **SQLite**: In-memory database for testing

## Contributing

1. Follow the existing code structure
2. Add tests for new functionality
3. Update documentation as needed
4. Ensure all tests pass before submitting

## License

This project is licensed under the MIT License.