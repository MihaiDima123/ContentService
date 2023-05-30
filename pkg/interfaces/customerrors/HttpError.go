package customerrors

type HTTPError interface {
	CustomError
	GetStatus() int
}
