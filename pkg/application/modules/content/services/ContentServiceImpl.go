package services

import (
	"contentservice/pkg/application/core/customErrors"
	"contentservice/pkg/application/core/customErrors/httpErrors"
	postDto "contentservice/pkg/application/dto/post"
	"contentservice/pkg/application/entity/post_entities"
	"contentservice/pkg/application/modules/content/interfaces"
	errorInterfaces "contentservice/pkg/interfaces/customerrors"
	"contentservice/pkg/interfaces/validator"
)

type ContentServiceImpl struct {
	contentRepository interfaces.ContentRepository
	postValidator     validator.Validator[post_entities.Post]
}

func (csi *ContentServiceImpl) Configure(repository interfaces.ContentRepository) {
	csi.contentRepository = repository
}

func New(contentRepository interfaces.ContentRepository,
	validator validator.Validator[post_entities.Post]) interfaces.ContentService {
	return &ContentServiceImpl{
		contentRepository: contentRepository,
		postValidator:     validator,
	}
}

func (csi *ContentServiceImpl) GetById(id int64) (*postDto.PostDTO, errorInterfaces.HTTPError) {
	// TODO: And this
	_, err := csi.contentRepository.GetById(id)
	if customErrors.Is(err, httpErrors.NotFoundErrorType) {
		return nil, httpErrors.HttpNotFoundError
	}

	if err != nil {
		return nil, httpErrors.HttpInternalServerError
	}

	// TODO: Change this
	return postDto.New(), httpErrors.HttpInternalServerError
}

func (csi *ContentServiceImpl) Create(post *post_entities.Post) (int64, errorInterfaces.HTTPError) {
	if validatorErrors := csi.postValidator.Validate(post); len(validatorErrors) != 0 {
		return 0, httpErrors.HttpBadRequestError
	}

	id, err := csi.contentRepository.Create(post)
	if customErrors.Is(err, httpErrors.InternalServerErrorType) {
		return 0, httpErrors.HttpInternalServerError
	}

	return id, nil
}
