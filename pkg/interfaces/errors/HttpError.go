package errors

type HTTPError interface {
	CustomError
	GetStatus() int
}
