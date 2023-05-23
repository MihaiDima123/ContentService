package validator

type Validator[T interface{}] interface {
	Validate(T) []ValidationError
}

type ValidationError struct {
	Message string
}
