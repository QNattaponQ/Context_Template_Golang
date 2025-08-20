package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	userRepo UserRepository
}

func NewUserHandler(userRepo UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

// CreateUser handles POST /users requests
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest

	// Parse request body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Basic validation
	if req.Name == "" || req.Email == "" || req.Age <= 0 || req.Age > 150 {
		http.Error(w, "Invalid input: name, email, and age (1-150) are required", http.StatusBadRequest)
		return
	}

	// Email format validation
	if _, err := mail.ParseAddress(req.Email); err != nil {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	// Create user model
	user := &User{
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
	}

	// Save to database
	if err := h.userRepo.Create(r.Context(), user); err != nil {
		if err.Error() == fmt.Sprintf("user with email %s already exists", req.Email) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   user.ToResponse(),
	})
}

// GetUser handles GET /users/{id} requests
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// Extract ID from URL parameters
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Parse ID
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Find user in database
	user, err := h.userRepo.FindByID(r.Context(), uint(id))
	if err != nil {
		if err.Error() == fmt.Sprintf("user with ID %d not found", id) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"data":   user.ToResponse(),
	})
}
