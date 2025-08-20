package errs

import (
	"fmt"
	"net/http"
)

// AppError represents application-specific errors
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

// NewAppError creates a new application error
func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// Common error constructors
func NewBadRequest(message string, err error) *AppError {
	return NewAppError(http.StatusBadRequest, message, err)
}

func NewNotFound(message string, err error) *AppError {
	return NewAppError(http.StatusNotFound, message, err)
}

func NewConflict(message string, err error) *AppError {
	return NewAppError(http.StatusConflict, message, err)
}

func NewInternalServerError(message string, err error) *AppError {
	return NewAppError(http.StatusInternalServerError, message, err)
}
