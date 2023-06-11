package httpErrors

import (
	"contentservice/pkg/application/core/customErrors"
	"contentservice/pkg/application/core/customErrors/utils"
	"contentservice/pkg/interfaces/customerrors"
	"net/http"
)

// NotFoundErrorType Error types
var NotFoundErrorType = &customErrors.ErrorType{
	Value:          1,
	DefaultMessage: "not found",
}

var InternalServerErrorType = &customErrors.ErrorType{
	Value:          2,
	DefaultMessage: "internal server error",
}

var BadRequestErrorType = &customErrors.ErrorType{
	Value:          3,
	DefaultMessage: "bad request",
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
