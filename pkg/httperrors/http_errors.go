package httperrors

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	errorBadRequest          = errors.New("bad request")
	errorNotFound            = errors.New("not found")
	errorUnauthorized        = errors.New("unauthorized")
	errorForbidden           = errors.New("forbidden")
	errorInternalServerError = errors.New("internal server error")
)

type RestErr interface {
	Error() string
	Status() int
	Causes() any
}

type restErr struct {
	status int
	error  string
	causes any
}

func (e restErr) Error() string {
	return fmt.Sprintf("status: %d - error: %s - causes: %v", e.status, e.error, e.causes)
}

func (e restErr) Status() int {
	return e.status
}

func (e restErr) Causes() any {
	return e.causes
}

func NewRestError(status int, err string, causes any) RestErr {
	return restErr{
		status: status,
		error:  err,
		causes: causes,
	}
}

func NewRestErrorWithMessage(status int, err string, causes any) RestErr {
	return restErr{
		status: status,
		error:  err,
		causes: causes,
	}
}

// BadRequest wraps causes with HTTP bad request error
func BadRequest(causes any) RestErr {
	return restErr{status: http.StatusBadRequest, error: errorBadRequest.Error(), causes: causes}
}

// Unauthorized wraps causes with HTTP unauthorized error
func Unauthorized(causes any) RestErr {
	return restErr{status: http.StatusUnauthorized, error: errorUnauthorized.Error(), causes: causes}
}

// NotFound wraps causes with HTTP not found error
func NotFound(causes any) RestErr {
	return restErr{status: http.StatusNotFound, error: errorNotFound.Error(), causes: causes}
}

// InternalServerError wraps causes with HTTP internal server error
func InternalServerError(causes any) RestErr {
	return restErr{status: http.StatusInternalServerError, error: errorInternalServerError.Error(), causes: causes}
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