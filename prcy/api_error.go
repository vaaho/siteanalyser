package prcy

import "fmt"

type ApiErrorType int

const (
	Unknown ApiErrorType = iota
	Limit
	Format
	Parse
	Network
)

type ApiError struct {
	Type    ApiErrorType
	Message string
}

func (e *ApiError) Error() string {
    return fmt.Sprintf("Error [Code: %d] %s", e.Type, e.Message)
}

func NewApiError(errorType ApiErrorType, message string) *ApiError {
	return &ApiError{Type:errorType, Message:message}
}