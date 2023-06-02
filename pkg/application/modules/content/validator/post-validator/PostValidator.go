package postValidator

import (
	"contentservice/pkg/application/core/customErrors/validation-errors"
	"contentservice/pkg/application/core/validator/instance"
	"contentservice/pkg/application/entity/post_entities"
	"contentservice/pkg/application/modules/content/validator"
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

	if validatorImpl.ValueLessThan(post.UserEmail, MinEmailLength) {
		errors = append(errors,
			validationErrors.GetTooShortError(EmailTooShortError))
	}

	return errors
}

func New() validator.Validator[post_entities.Post] {
	return new(PostValidator)
}
