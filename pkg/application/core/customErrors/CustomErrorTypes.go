package customErrors

import (
	"contentservice/pkg/interfaces/customerrors"
)

type ErrorType struct {
	Value          int8
	DefaultMessage string
}

// Is checks if an error is of type
func Is(httpError customerrors.CustomError, et *ErrorType) bool {
	return httpError.GetErrorType() == et.Value
}
