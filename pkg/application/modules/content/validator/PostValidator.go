package validator

import (
	"contentservice/pkg/application/core/customErrors/validation-errors"
	"contentservice/pkg/application/entity/post_entities"
	"contentservice/pkg/interfaces/customerrors"
	"contentservice/pkg/interfaces/validator"
)

const (
	EmailTooShortError = "Email too short"
)

const (
	MinEmailLength = 3
)

type PostValidator struct {
	validator.Validator[post_entities.Post]
}

func (pv *PostValidator) Validate(post post_entities.Post) []customerrors.ValidationError {
	var errors []customerrors.ValidationError

	if valueLessThan(post.UserEmail, MinEmailLength) {
		errors = append(errors, validation_errors.GetTooShortError("Email"))
	}

	return errors
}

func NewPostValidator() *PostValidator {
	return new(PostValidator)
}
