package errors

import "net/http"

type BusinessError struct {
	Message    string
	StatusCode int
}

func (e *BusinessError) Error() string {
	return e.Message
}

func NewBusinessError(message string, statusCode ...int) error {
	code := http.StatusBadRequest
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	return &BusinessError{
		Message:    message,
		StatusCode: code,
	}
}

type InfrastructureError struct {
	Message    string
	StatusCode int
}

func (e *InfrastructureError) Error() string {
	return e.Message
}

func NewInfrastructureError(message string, statusCode ...int) error {
	code := http.StatusInternalServerError
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	return &InfrastructureError{
		Message:    message,
		StatusCode: code,
	}
}

// Predefined errors
var (
	ErrorNotFound       = NewBusinessError("not found", http.StatusNotFound)
	ErrorInfrastructure = NewInfrastructureError("infrastructure error")
	ErrorInternal       = NewInfrastructureError("internal server error")
)
