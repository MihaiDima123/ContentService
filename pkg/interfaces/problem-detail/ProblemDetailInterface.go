package problemDetail

type ProblemDetail interface {
	Title(title string) ProblemDetail
	Status(status int) ProblemDetail
	Instance(instance string) ProblemDetail
	Detail(detail string) ProblemDetail
	Properties(properties any) ProblemDetail
}
