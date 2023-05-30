package interfaces

type CustomError interface {
	error
	GetErrorType() int8
}
