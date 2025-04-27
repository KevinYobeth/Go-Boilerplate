package errors

import (
	"net/http"
)

var ErrorMap = map[ErrorType]int{
	ErrorTypeIncorrectInput:  http.StatusBadRequest,
	ErrorTypeNotFound:        http.StatusNotFound,
	ErrorTypeUnknown:         http.StatusInternalServerError,
	ErrorTypeUnauthenticated: http.StatusUnauthorized,
	ErrorTypeUnauthorized:    http.StatusForbidden,
}
