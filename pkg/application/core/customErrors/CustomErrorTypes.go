package customErrors

import (
	"contentservice/pkg/interfaces/errors"
)

type errorType struct {
	Value          int8
	DefaultMessage string
}

type CustomHttpError struct {
	error
	ErrorType int8
	Status    int
}

type CustomDbError struct {
	error
	ErrorType int8
}

// GetErrorType Db error
func (cde *CustomDbError) GetErrorType() int8 {
	return cde.ErrorType
}

func (che *CustomHttpError) GetStatus() int {
	return che.Status
}

// GetErrorType Http error
func (che *CustomHttpError) GetErrorType() int8 {
	return che.ErrorType
}

// Is checks if an error is of type
func Is(httpError errors.CustomError, et errorType) bool {
	return httpError.GetErrorType() == et.Value
}
