package validator

import (
	"contentservice/pkg/application/entity/post_entities"
	"contentservice/pkg/interfaces/validator"
)

const (
	EmailTooShortError = "Email too short"
)

type PostValidator struct {
	validator.Validator[post_entities.Post]
}

func (pv *PostValidator) Validate(post post_entities.Post) []validator.ValidationError {
	var errors []validator.ValidationError

	if len(post.UserEmail) <= 3 {
		errors = append(errors, validator.ValidationError{Message: EmailTooShortError})
	}

	return errors
}

func NewPostValidator() *PostValidator {
	return new(PostValidator)
}
