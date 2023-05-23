package services

import (
	"contentservice/pkg/application/core/customErrors"
	"contentservice/pkg/application/entity/post_entities"
	"contentservice/pkg/application/modules/content/interfaces"
	"contentservice/pkg/application/modules/content/validator"
	"errors"
)

type ContentServiceImpl struct {
	contentRepository interfaces.ContentRepository
	postValidator     *validator.PostValidator
}

func (csi *ContentServiceImpl) Configure(repository interfaces.ContentRepository) {
	csi.contentRepository = repository
}

func New(contentRepository interfaces.ContentRepository,
	validator *validator.PostValidator) interfaces.ContentService {
	return &ContentServiceImpl{
		contentRepository: contentRepository,
		postValidator:     validator,
	}
}

func (csi *ContentServiceImpl) GetById(id int64) (*post_entities.Post, customErrors.HTTPError) {
	post, err := csi.contentRepository.GetById(id)
	if errors.Is(err, customErrors.ResourceNotFoundError) {
		return post, customErrors.NotFoundError
	}

	if err != nil {
		return nil, customErrors.InternalServerError
	}

	return post, customErrors.InternalServerError
}

func (csi *ContentServiceImpl) Create(post post_entities.Post) (int64, customErrors.HTTPError) {
	if validatorErrors := csi.postValidator.Validate(post); len(validatorErrors) != 0 {
		return 0, customErrors.BadRequestError
	}

	id, err := csi.contentRepository.Create(post)

	if errors.Is(err, customErrors.ResourceNotCreatedError) || err != nil {
		return 0, customErrors.InternalServerError
	}

	return id, nil
}
