package customErrors

import (
	"errors"
	"net/http"
)

var NotFoundError = &CustomHttpError{
	error:  errors.New("not found"),
	Status: http.StatusNotFound,
}

var InternalServerError = &CustomHttpError{
	error:  errors.New("not created"),
	Status: http.StatusInternalServerError,
}
