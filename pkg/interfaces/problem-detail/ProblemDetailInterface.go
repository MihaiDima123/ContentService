package problemDetail

type ProblemDetail[T any] interface {
	Title(title string) ProblemDetail[T]
	Status(status int) ProblemDetail[T]
	Instance(instance string) ProblemDetail[T]
	Detail(detail string) ProblemDetail[T]
	Properties(properties []T) ProblemDetail[T]
}
