package validator

import "contentservice/pkg/interfaces/customerrors"

// Validator TODO: Make a chain-like validator
type Validator[T any] interface {
	Validate(T) []customerrors.ValidationError
}
