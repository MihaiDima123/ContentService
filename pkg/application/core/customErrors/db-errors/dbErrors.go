package dbErrors

import (
	"contentservice/pkg/application/core/customErrors"
	"contentservice/pkg/application/core/customErrors/utils"
	"contentservice/pkg/interfaces/customerrors"
)

var DbResourceNotFoundType = &customErrors.ErrorType{
	Value:          1,
	DefaultMessage: "resource not found",
}

var DbResourceNotCreatedType = &customErrors.ErrorType{
	Value:          2,
	DefaultMessage: "resource not created",
}

var DbResourceNotFetchedType = &customErrors.ErrorType{
	Value:          3,
	DefaultMessage: "resource coould not be fetched",
}

// DbNotFoundError Constants
var DbNotFoundError = GetResourceNotFoundError()
var DbNotCreatedError = GetResourceNotCreatedError()
var DbNotFetchedError = GetResourceNotFetchedError()

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

var GetResourceNotFetchedError = func(params ...string) customerrors.DbError {
	return &CustomDbError{
		error:     utils.GetErrorFromString(DbResourceNotFetchedType.DefaultMessage, params...),
		ErrorType: DbResourceNotFetchedType.Value,
	}
}
