package validationErrors

import (
	"contentservice/pkg/application/core/customErrors"
	"contentservice/pkg/application/core/customErrors/utils"
	"contentservice/pkg/interfaces/customerrors"
	"contentservice/pkg/interfaces/validator"
)

var TooShortErrorType = customErrors.ErrorType{
	Value:          1,
	DefaultMessage: "Too short",
}

var CouldNotParseErrorType = customErrors.ErrorType{
	Value:          2,
	DefaultMessage: "Could not parse",
}

func GetTooShortError(instance validator.ValidationInstance) customerrors.ValidationError {
	return &CustomValidationError{
		error:     utils.GetErrorFromString(TooShortErrorType.DefaultMessage, instance.Message()),
		Source:    instance.Instance(),
		ErrorType: TooShortErrorType.Value,
	}
}

func GetParseError(instance validator.ValidationInstance) customerrors.ValidationError {
	return &CustomValidationError{
		error:     utils.GetErrorFromString(CouldNotParseErrorType.DefaultMessage, instance.Message()),
		Source:    instance.Instance(),
		ErrorType: CouldNotParseErrorType.Value,
	}
}
