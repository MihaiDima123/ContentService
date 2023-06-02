package validator

import (
	"contentservice/pkg/application/core/customErrors/validation-errors"
	"contentservice/pkg/application/entity/post_entities"
	"contentservice/pkg/application/modules/content/validator/instance"
	"contentservice/pkg/interfaces/customerrors"
	"contentservice/pkg/interfaces/validator"
)

// instances

var EmailTooShortError = validationInstance.New("Email", "too short")

// constants

const MinEmailLength = 3

type PostValidator struct {
	validator.Validator[post_entities.Post]
}

func (pv *PostValidator) Validate(post post_entities.Post) []customerrors.ValidationError {
	var errors []customerrors.ValidationError

	if valueLessThan(post.UserEmail, MinEmailLength) {
		errors = append(errors,
			validationErrors.GetTooShortError(EmailTooShortError))
	}

	return errors
}

func NewPostValidator() *PostValidator {
	return new(PostValidator)
}
