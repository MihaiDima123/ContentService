package dbErrors

type CustomDbError struct {
	error
	ErrorType int8
}

func (cde *CustomDbError) GetErrorType() int8 {
	return cde.ErrorType
}
