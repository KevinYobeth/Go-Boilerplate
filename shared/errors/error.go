package errors

import (
	"errors"
	"fmt"

	"github.com/ztrue/tracerr"
)

type ErrorType string

var (
	ErrorTypeUnknown        = ErrorType("unknown")
	ErrorTypeDatabase       = ErrorType("database")
	ErrorTypeIncorrectInput = ErrorType("incorrect-input")
	ErrorTypeNotFound       = ErrorType("not-found")
)

type GenericError struct {
	Err     error
	Type    ErrorType
	Message string
}

func (e GenericError) Unwrap() error {
	return e.Err
}

func (e GenericError) Error() string {
	return e.Message
}

func (e GenericError) ErrorType() ErrorType {
	return e.Type
}

func NewGenericError(err error, message string) GenericError {
	return GenericError{
		Err:     err,
		Type:    ErrorTypeUnknown,
		Message: message,
	}
}

func NewInitializationError(err error, service string) GenericError {
	return GenericError{
		Err:     err,
		Type:    ErrorTypeUnknown,
		Message: fmt.Sprintf("failed to initialize %s service: %s", service, err),
	}
}

func NewNotFoundError(err error, resource string) GenericError {
	return GenericError{
		Err:     err,
		Type:    ErrorTypeNotFound,
		Message: fmt.Sprintf("%s not found", resource),
	}
}

func NewIncorrectInputError(err error, message string) GenericError {
	return GenericError{
		Err:     err,
		Type:    ErrorTypeIncorrectInput,
		Message: message,
	}
}

func GetGenericError(err error) GenericError {
	if err == nil {
		return NewGenericError(err, "")
	}

	genericErr, ok := err.(GenericError)
	if ok {
		return genericErr
	}

	return GetGenericError(errors.Unwrap(err))
}

func GetTracerrErr(err error) tracerr.Error {
	if err == nil {
		return nil
	}

	unwrappedErr, ok := err.(tracerr.Error)
	if ok {
		return unwrappedErr
	}

	return GetTracerrErr(errors.Unwrap(err))
}
