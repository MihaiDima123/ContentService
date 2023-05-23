package customErrors

import "errors"

const (
	ResourceNotFound   int8 = 1
	ResourceNotCreated int8 = 2
)

var ResourceNotFoundError = &CustomDbError{
	error:     errors.New("resource not found"),
	ErrorType: ResourceNotFound,
}

var ResourceNotCreatedError = &CustomDbError{
	error:     errors.New("resource not created"),
	ErrorType: ResourceNotCreated,
}
