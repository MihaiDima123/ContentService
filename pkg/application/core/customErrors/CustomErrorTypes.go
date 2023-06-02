package customErrors

import (
	"contentservice/pkg/interfaces/customerrors"
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

type CustomValidationError struct {
	error
	ErrorType int8
	Source    string
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

// GetErrorType ValidationError
func (cve *CustomValidationError) GetErrorType() int8 {
	return cve.ErrorType
}

func (cve *CustomValidationError) GetSource() string {
	return cve.Source
}

// Is checks if an error is of type
func Is(httpError customerrors.CustomError, et *errorType) bool {
	return httpError.GetErrorType() == et.Value
}
