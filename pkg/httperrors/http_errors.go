package httperrors

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	// ErrBadRequest represents bad request error
	ErrBadRequest = errors.New("bad request")
	// ErrNotFound represents not found error
	ErrNotFound = errors.New("not found")
	// ErrUnauthorized represents unauthorized error
	ErrUnauthorized = errors.New("unauthorized")
	// ErrForbidden represents forbidden error
	ErrForbidden = errors.New("forbidden")
	// ErrInternalServerError represents internal server error
	ErrInternalServerError = errors.New("internal server error")
)

// RestErr represents the http error with status code, message and its wrapped cause
type RestErr interface {
	Error() string
	Status() int
	Causes() any
}

type restError struct {
	status int
	error  string
	causes any
}

// Error fills error.
func (e restError) Error() string {
	return fmt.Sprintf("status: %d - error: %s - causes: %v", e.status, e.error, e.causes)
}

func (e restError) Status() int {
	return e.status
}

func (e restError) Causes() any {
	return e.causes
}

// NewRestError creates RestErr with status, message and wrapped cause
func NewRestError(status int, err string, causes any) RestErr {
	return restError{
		status: status,
		error:  err,
		causes: causes,
	}
}

// BadRequest wraps causes with HTTP bad request error
func BadRequest(causes any) RestErr {
	return restError{status: http.StatusBadRequest, error: ErrBadRequest.Error(), causes: causes}
}

// Unauthorized wraps causes with HTTP unauthorized error
func Unauthorized(causes any) RestErr {
	return restError{status: http.StatusUnauthorized, error: ErrUnauthorized.Error(), causes: causes}
}

// NotFound wraps causes with HTTP not found error
func NotFound(causes any) RestErr {
	return restError{status: http.StatusNotFound, error: ErrNotFound.Error(), causes: causes}
}

// InternalServerError wraps causes with HTTP internal server error
func InternalServerError(causes any) RestErr {
	return restError{status: http.StatusInternalServerError, error: ErrInternalServerError.Error(), causes: causes}
}

// ParseError defines which kind of HTTP error err is
func ParseError(err error) RestErr {
	switch {
	default:
		if restErr, ok := err.(RestErr); ok {
			return restErr
		}
		return InternalServerError(err)
	}
}

// ErrorResponse returns error status and body
func ErrorResponse(err error) (int, gin.H) {
	e := ParseError(err)

	return e.Status(), gin.H{"msg": e.Error()}
}
