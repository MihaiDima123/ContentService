package customErrors

import (
	"contentservice/pkg/application/core/customErrors/utils"
	"contentservice/pkg/interfaces/customerrors"
)

var DbResourceNotFoundType = errorType{
	Value:          1,
	DefaultMessage: "resource not found",
}

var DbResourceNotCreatedType = errorType{
	Value:          2,
	DefaultMessage: "resource not created",
}

// DbNotFoundError Constants
var DbNotFoundError = GetResourceNotFoundError()
var DbNotCreatedError = GetResourceNotCreatedError()

var GetResourceNotFoundError = func(params ...string) customerrors.DbError {
	return &CustomDbError{
		error:     utils.GetErrorFromString(DbResourceNotFoundType.DefaultMessage, params...),
		ErrorType: DbResourceNotFoundType.Value,
	}
}

var GetResourceNotCreatedError = func(params ...string) customerrors.DbError {
	return &CustomDbError{
		error:     utils.GetErrorFromString(DbResourceNotCreatedType.DefaultMessage, params...),
		ErrorType: DbResourceNotCreatedType.Value,
	}
}
