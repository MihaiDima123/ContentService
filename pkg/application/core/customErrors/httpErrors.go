package customErrors

import (
	"contentservice/pkg/application/core/customErrors/interfaces"
	"contentservice/pkg/application/core/customErrors/utils"
	"net/http"
)

var NotFoundErrorType = errorType{
	Value:          1,
	DefaultMessage: "not found",
}

var InternalServerErrorType = errorType{
	Value:          2,
	DefaultMessage: "not created",
}

var BadRequestErrorType = errorType{
	Value:          3,
	DefaultMessage: "internal server error",
}

var NotFoundError = func(params ...string) interfaces.HTTPError {
	return &CustomHttpError{
		error:     utils.GetErrorFromString(NotFoundErrorType.DefaultMessage, params...),
		ErrorType: NotFoundErrorType.Value,
		Status:    http.StatusNotFound,
	}
}

var InternalServerError = func(params ...string) interfaces.HTTPError {
	return &CustomHttpError{
		error:     utils.GetErrorFromString(InternalServerErrorType.DefaultMessage, params...),
		ErrorType: InternalServerErrorType.Value,
		Status:    http.StatusInternalServerError,
	}
}

var BadRequestError = func(params ...string) interfaces.HTTPError {
	return &CustomHttpError{
		error:     utils.GetErrorFromString(BadRequestErrorType.DefaultMessage, params...),
		ErrorType: BadRequestErrorType.Value,
		Status:    http.StatusBadRequest,
	}
}
