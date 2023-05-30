package errors

type CustomError interface {
	error
	GetErrorType() int8
}
