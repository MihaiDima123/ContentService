package interfaces

type CustomError interface {
	error
	GetErrorType() int8
}

type DbError interface {
	CustomError
}

type HTTPError interface {
	CustomError
	GetStatus() int
}
