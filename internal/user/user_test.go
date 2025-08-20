package user

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		t.Fatal(err)
	}
	return db
}

func setupTestHandler(t *testing.T) (*UserHandler, *gorm.DB) {
	db := setupTestDB(t)
	userRepo := NewUserRepository(db)
	userHandler := NewUserHandler(userRepo)
	return userHandler, db
}

func TestUserRepository_CreateAndFind(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)

	ctx := context.Background()
	user := &User{Name: "Alice", Email: "alice@example.com", Age: 30}

	// Test Create
	err := repo.Create(ctx, user)
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)

	// Test FindByID
	found, err := repo.FindByID(ctx, user.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Alice", found.Name)
	assert.Equal(t, "alice@example.com", found.Email)
	assert.Equal(t, 30, found.Age)
}

func TestUserRepository_DuplicateEmail(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)

	ctx := context.Background()
	user1 := &User{Name: "Alice", Email: "alice@example.com", Age: 30}
	user2 := &User{Name: "Bob", Email: "alice@example.com", Age: 25}

	// Create first user
	err := repo.Create(ctx, user1)
	assert.NoError(t, err)

	// Try to create second user with same email
	err = repo.Create(ctx, user2)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already exists")
}

func TestUserRepository_UserNotFound(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)

	ctx := context.Background()
	_, err := repo.FindByID(ctx, 999)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not found")
}

func TestUserHandler_CreateUser(t *testing.T) {
	tests := []struct {
		name           string
		input          CreateUserRequest
		expectedStatus int
		expectedError  bool
	}{
		{
			name: "valid user",
			input: CreateUserRequest{
				Name:  "John Doe",
				Email: "john@example.com",
				Age:   30,
			},
			expectedStatus: http.StatusCreated,
			expectedError:  false,
		},
		{
			name: "invalid email",
			input: CreateUserRequest{
				Name:  "John Doe",
				Email: "invalid-email",
				Age:   30,
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
		{
			name: "invalid age - too young",
			input: CreateUserRequest{
				Name:  "John Doe",
				Email: "john@example.com",
				Age:   0,
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
		{
			name: "invalid age - too old",
			input: CreateUserRequest{
				Name:  "John Doe",
				Email: "john@example.com",
				Age:   151,
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
		{
			name: "missing name",
			input: CreateUserRequest{
				Name:  "",
				Email: "john@example.com",
				Age:   30,
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler, _ := setupTestHandler(t)

			// Create request body
			body, _ := json.Marshal(tt.input)
			req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			// Create response recorder
			w := httptest.NewRecorder()

			// Call handler
			handler.CreateUser(w, req)

			// Assert response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if !tt.expectedError {
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, "success", response["status"])
			}
		})
	}
}

func TestUserHandler_GetUser(t *testing.T) {
	tests := []struct {
		name           string
		userID         string
		setupUser      bool
		expectedStatus int
		expectedError  bool
	}{
		{
			name:           "valid user ID",
			userID:         "1",
			setupUser:      true,
			expectedStatus: http.StatusOK,
			expectedError:  false,
		},
		{
			name:           "user not found",
			userID:         "999",
			setupUser:      false,
			expectedStatus: http.StatusNotFound,
			expectedError:  true,
		},
		{
			name:           "invalid user ID",
			userID:         "invalid",
			setupUser:      false,
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
		{
			name:           "missing user ID",
			userID:         "",
			setupUser:      false,
			expectedStatus: http.StatusBadRequest,
			expectedError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler, db := setupTestHandler(t)

			// Setup test user if needed
			if tt.setupUser {
				user := &User{Name: "Test User", Email: "test@example.com", Age: 25}
				db.Create(user)
			}

			// Create request
			req := httptest.NewRequest("GET", "/users/"+tt.userID, nil)

			// Setup Chi context for URL parameters
			rctx := chi.NewRouteContext()
			if tt.userID != "" {
				rctx.URLParams.Add("id", tt.userID)
			}
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			// Create response recorder
			w := httptest.NewRecorder()

			// Call handler
			handler.GetUser(w, req)

			// Assert response
			assert.Equal(t, tt.expectedStatus, w.Code)

			if !tt.expectedError && tt.setupUser {
				var response map[string]interface{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, "success", response["status"])
			}
		})
	}
}

func TestUserHandler_DuplicateEmailConflict(t *testing.T) {
	handler, db := setupTestHandler(t)

	// Create first user
	user1 := &User{Name: "Alice", Email: "alice@example.com", Age: 30}
	db.Create(user1)

	// Try to create second user with same email
	reqBody := CreateUserRequest{
		Name:  "Bob",
		Email: "alice@example.com",
		Age:   25,
	}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler.CreateUser(w, req)

	// Should return conflict status
	assert.Equal(t, http.StatusConflict, w.Code)
}
