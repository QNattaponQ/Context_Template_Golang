# PRP: REST API for User Management (Golang)

## Overview
Implement a REST API in Golang with endpoints for user creation and retrieval, using PostgreSQL for storage and GORM as the ORM. Follow project patterns and conventions as demonstrated in provided examples.

---

## Feature Requirements
- **POST `/users`**: Create a new user. Accepts JSON payload with `name`, `email`, and `age`.
- **GET `/users/{id}`**: Retrieve user details by ID.
- Data must be persisted in PostgreSQL using GORM.

---

## Implementation Steps
1. **Define User Model**
   - Fields: `id`, `name`, `email`, `age`, `created_at`
   - Use GORM struct tags for mapping and validation
2. **Set Up Database Connection**
   - Use GORM to connect to PostgreSQL
   - Configure connection in a reusable pattern (see `examples/repository.go`)
3. **Create Repository Layer**
   - Implement CRUD operations for users
   - Follow repository pattern from `examples/repository.go`
4. **Set Up HTTP Server**
   - Use chi router for routing (see `examples/server.go`)
   - Register `/users` endpoints
5. **Implement Handlers**
   - POST: Validate input, create user, handle errors
   - GET: Fetch user by ID, handle not found
6. **Testing**
   - Write unit tests for repository and handlers
   - Use table-driven tests as in `examples/tests/test_example.go`
7. **Validation & Error Handling**
   - Ensure proper error responses for invalid input, DB errors, and not found
8. **Documentation**
   - Document endpoints, request/response formats, and error codes

---

## Example Patterns to Follow
- **HTTP Server Setup**: `examples/server.go`
- **Repository Pattern**: `examples/repository.go`
- **Unit Testing**: `examples/tests/test_example.go`

---

## Database Schema
```sql
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  email TEXT UNIQUE NOT NULL,
  age INT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);
```

---

## References
- [GORM Documentation](https://gorm.io/docs/)
- [Chi Router](https://github.com/go-chi/chi)

---

## Validation Gates
- All endpoints must pass unit tests
- Database operations must be covered by tests
- Error handling must be robust and follow conventions

---

## Confidence Score
- 9/10 (based on provided context and examples)

---

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

Generated on August 20, 2025.
