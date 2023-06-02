package validationErrors

import (
	"contentservice/pkg/application/core/customErrors"
	"contentservice/pkg/application/core/customErrors/utils"
	"contentservice/pkg/interfaces/customerrors"
)

var TooShortErrorType = customErrors.ErrorType{
	Value:          1,
	DefaultMessage: "Too short",
}

func GetTooShortError(source string, reasons ...string) customerrors.ValidationError {
	return &CustomValidationError{
		error:     utils.GetErrorFromString(TooShortErrorType.DefaultMessage, reasons...),
		Source:    source,
		ErrorType: TooShortErrorType.Value,
	}
}
