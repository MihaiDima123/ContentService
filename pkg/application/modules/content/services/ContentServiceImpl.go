package services

import (
	"contentservice/pkg/application/core/customErrors"
	"contentservice/pkg/application/entity/post_entities"
	"contentservice/pkg/application/modules/content/interfaces"
	"contentservice/pkg/application/modules/content/validator"
	errorInterfaces "contentservice/pkg/interfaces/customerrors"
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

func (csi *ContentServiceImpl) GetById(id int64) (*post_entities.Post, errorInterfaces.HTTPError) {
	post, err := csi.contentRepository.GetById(id)
	if customErrors.Is(err, customErrors.NotFoundErrorType) {
		return post, customErrors.HttpNotFoundError
	}

	if err != nil {
		return nil, customErrors.HttpInternalServerError
	}

	return post, customErrors.HttpInternalServerError
}

func (csi *ContentServiceImpl) Create(post post_entities.Post) (int64, errorInterfaces.HTTPError) {
	if validatorErrors := csi.postValidator.Validate(post); len(validatorErrors) != 0 {
		return 0, customErrors.HttpBadRequestError
	}

	id, err := csi.contentRepository.Create(post)
	if customErrors.Is(err, customErrors.InternalServerErrorType) {
		return 0, customErrors.HttpInternalServerError
	}

	return id, nil
}
