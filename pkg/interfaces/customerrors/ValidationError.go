package customerrors

type ValidationError interface {
	CustomError
	GetSource() string
	GetTitle() string
}
