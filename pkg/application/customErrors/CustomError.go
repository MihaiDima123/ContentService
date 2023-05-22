package customErrors

type HTTPError interface {
	error
	GetStatus() int
}

type DbError interface {
	error
	GetErrorType() int8
}

type CustomHttpError struct {
	error
	Status int
}

func (che *CustomHttpError) GetStatus() int {
	return che.Status
}

type CustomDbError struct {
	error
	ErrorType int8
}

func (cde *CustomDbError) GetErrorType() int8 {
	return cde.ErrorType
}
