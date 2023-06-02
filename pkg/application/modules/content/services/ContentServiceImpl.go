package services

import (
	"contentservice/pkg/application/core/customErrors"
	"contentservice/pkg/application/core/customErrors/httpErrors"
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
	if customErrors.Is(err, httpErrors.NotFoundErrorType) {
		return post, httpErrors.HttpNotFoundError
	}

	if err != nil {
		return nil, httpErrors.HttpInternalServerError
	}

	return post, httpErrors.HttpInternalServerError
}

func (csi *ContentServiceImpl) Create(post post_entities.Post) (int64, errorInterfaces.HTTPError) {
	if validatorErrors := csi.postValidator.Validate(post); len(validatorErrors) != 0 {
		return 0, httpErrors.HttpBadRequestError
	}

	id, err := csi.contentRepository.Create(post)
	if customErrors.Is(err, httpErrors.InternalServerErrorType) {
		return 0, httpErrors.HttpInternalServerError
	}

	return id, nil
}
