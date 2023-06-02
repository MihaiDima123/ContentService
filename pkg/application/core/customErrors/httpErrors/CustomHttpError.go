package httpErrors

type CustomHttpError struct {
	error
	ErrorType int8
	Status    int
}

func (che *CustomHttpError) GetStatus() int {
	return che.Status
}

func (che *CustomHttpError) GetErrorType() int8 {
	return che.ErrorType
}
