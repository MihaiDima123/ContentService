package customErrors

import (
	"contentservice/pkg/application/core/customErrors/interfaces"
	"contentservice/pkg/application/core/customErrors/utils"
)

var DbResourceNotFoundType = errorType{
	Value:          1,
	DefaultMessage: "resource not found",
}

var DbResourceNotCreatedType = errorType{
	Value:          2,
	DefaultMessage: "resource not created",
}

var ResourceNotFoundError = func(params ...string) interfaces.DbError {
	return &CustomDbError{
		error:     utils.GetErrorFromString(DbResourceNotFoundType.DefaultMessage, params...),
		ErrorType: DbResourceNotFoundType.Value,
	}
}

var ResourceNotCreatedError = func(params ...string) interfaces.DbError {
	return &CustomDbError{
		error:     utils.GetErrorFromString(DbResourceNotCreatedType.DefaultMessage, params...),
		ErrorType: DbResourceNotCreatedType.Value,
	}
}
