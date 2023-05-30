package customerrors

type CustomError interface {
	error
	GetErrorType() int8
}
