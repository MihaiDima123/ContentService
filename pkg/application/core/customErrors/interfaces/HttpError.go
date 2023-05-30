package interfaces

type HTTPError interface {
	CustomError
	GetStatus() int
}
