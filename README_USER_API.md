# User Management REST API

A Go-based REST API for user management operations using PostgreSQL and GORM.

## Features

- **POST** `/users` - Create new users
- **GET** `/users/{id}` - Retrieve user by ID
- PostgreSQL database with GORM ORM
- Comprehensive input validation
- Proper error handling and HTTP status codes
- Table-driven tests with SQLite for testing

## Prerequisites

- Go 1.21 or higher
- Docker and Docker Compose (for PostgreSQL)
- PostgreSQL (if not using Docker)

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
# Copy and modify as needed
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=password
export DB_NAME=userdb
export SERVER_PORT=8080
```

### 4. Run Database Migration

```bash
# Connect to PostgreSQL and run the migration
psql -h localhost -U postgres -d userdb -f migrations/001_create_users_table.sql
```

### 5. Start the Server

```bash
go run main.go
```

The server will start on port 8080 (or the port specified in SERVER_PORT).

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

**Response:**
```json
{
  "status": "success",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "age": 30,
    "created_at": "2024-01-01T12:00:00Z"
  }
}
```

### Get User by ID

```bash
curl http://localhost:8080/users/1
```

**Response:**
```json
{
  "status": "success",
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "age": 30,
    "created_at": "2024-01-01T12:00:00Z"
  }
}
```

### Health Check

```bash
curl http://localhost:8080/health
```

**Response:**
```json
{"status":"ok"}
```

## Validation Rules

- **Name**: Required, non-empty string
- **Email**: Required, valid email format, must be unique
- **Age**: Required, integer between 1 and 150

## Error Handling

The API returns appropriate HTTP status codes:

- `200 OK` - Successful GET request
- `201 Created` - Successful user creation
- `400 Bad Request` - Invalid input data
- `404 Not Found` - User not found
- `409 Conflict` - Email already exists
- `500 Internal Server Error` - Server/database error

## Testing

Run the comprehensive test suite:

```bash
# Run all tests
go test ./examples/test/ -v

# Run specific test
go test ./examples/test/ -run TestUserHandler_CreateUser -v

# Run with coverage
go test ./examples/test/ -cover
```

## Project Structure

```
.
├── config/
│   └── config.go              # Configuration management
├── examples/
│   ├── database/
│   │   └── database.go        # Database connection
│   ├── handlers/
│   │   └── user_handler.go    # HTTP handlers
│   ├── models/
│   │   └── user.go            # Data models
│   ├── repository/
│   │   └── repository.go      # Data access layer
│   ├── server/
│   │   └── server.go          # HTTP server setup
│   └── test/
│       └── test_user.go       # Comprehensive tests
├── migrations/
│   └── 001_create_users_table.sql  # Database schema
├── docker-compose.yml          # PostgreSQL setup
├── go.mod                      # Go module dependencies
├── main.go                     # Application entry point
└── README_USER_API.md          # This file
```

## Development

### Adding New Features

1. Create models in `examples/models/`
2. Add repository methods in `examples/repository/`
3. Create handlers in `examples/handlers/`
4. Add routes in `examples/server/server.go`
5. Write tests in `examples/test/`

### Database Changes

1. Create new migration files in `migrations/`
2. Update models as needed
3. Test with both PostgreSQL and SQLite (for tests)

## Troubleshooting

### Common Issues

1. **Database Connection Failed**
   - Check PostgreSQL is running
   - Verify connection parameters in environment variables
   - Ensure database 'userdb' exists

2. **Port Already in Use**
   - Change SERVER_PORT environment variable
   - Check if another service is using port 8080

3. **Test Failures**
   - Ensure all dependencies are installed: `go mod tidy`
   - Check Go version compatibility
   - Verify test database setup

### Logs

The server provides detailed logging for debugging:
- Database connection status
- HTTP request/response details
- Error stack traces

## Contributing

1. Follow existing code patterns
2. Add tests for new functionality
3. Update documentation as needed
4. Ensure all tests pass before submitting

## License

This project is licensed under the MIT License.
