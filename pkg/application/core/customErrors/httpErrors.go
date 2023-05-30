package customErrors

import (
	"contentservice/pkg/application/core/customErrors/utils"
	"contentservice/pkg/interfaces/customerrors"
	"net/http"
)

// NotFoundErrorType Error types
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

// HttpNotFoundError Constants
var HttpNotFoundError = GetNotFoundError()
var HttpBadRequestError = GetBadRequestError()
var HttpInternalServerError = GetInternalServerError()

var GetNotFoundError = func(params ...string) customerrors.HTTPError {
	return &CustomHttpError{
		error:     utils.GetErrorFromString(NotFoundErrorType.DefaultMessage, params...),
		ErrorType: NotFoundErrorType.Value,
		Status:    http.StatusNotFound,
	}
}

var GetInternalServerError = func(params ...string) customerrors.HTTPError {
	return &CustomHttpError{
		error:     utils.GetErrorFromString(InternalServerErrorType.DefaultMessage, params...),
		ErrorType: InternalServerErrorType.Value,
		Status:    http.StatusInternalServerError,
	}
}

var GetBadRequestError = func(params ...string) customerrors.HTTPError {
	return &CustomHttpError{
		error:     utils.GetErrorFromString(BadRequestErrorType.DefaultMessage, params...),
		ErrorType: BadRequestErrorType.Value,
		Status:    http.StatusBadRequest,
	}
}
