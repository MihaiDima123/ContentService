package validator

import "contentservice/pkg/interfaces/customerrors"

type Validator[T any] interface {
	Validate(T) []customerrors.ValidationError
}
