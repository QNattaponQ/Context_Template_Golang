package responses

import (
	"encoding/json"
	"net/http"

	"yourproject/internal/shared/errs"
)

// Response represents a standard API response
type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// Success sends a successful response
func Success(w http.ResponseWriter, statusCode int, data interface{}) {
	response := Response{
		Status: "success",
		Data:   data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

// Error sends an error response
func Error(w http.ResponseWriter, appErr *errs.AppError) {
	response := Response{
		Status: "error",
		Error:  appErr.Message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(appErr.Code)
	json.NewEncoder(w).Encode(response)
}

// BadRequest sends a 400 Bad Request response
func BadRequest(w http.ResponseWriter, message string) {
	appErr := errs.NewBadRequest(message, nil)
	Error(w, appErr)
}

// NotFound sends a 404 Not Found response
func NotFound(w http.ResponseWriter, message string) {
	appErr := errs.NewNotFound(message, nil)
	Error(w, appErr)
}

// Conflict sends a 409 Conflict response
func Conflict(w http.ResponseWriter, message string) {
	appErr := errs.NewConflict(message, nil)
	Error(w, appErr)
}

// InternalServerError sends a 500 Internal Server Error response
func InternalServerError(w http.ResponseWriter, message string) {
	appErr := errs.NewInternalServerError(message, nil)
	Error(w, appErr)
}
