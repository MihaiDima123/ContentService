package validator

// Validator TODO: Make a chain-like validator
type Validator[T any] interface {
	Validate(T) []ValidationError
}

type ValidationError struct {
	Message string
}
