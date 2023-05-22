package services

import (
	"contentservice/pkg/application/customErrors"
	"contentservice/pkg/application/entity/post_entities"
	"contentservice/pkg/application/modules/content/interfaces"
	"errors"
)

type ContentServiceImpl struct {
	contentRepository interfaces.ContentRepository
}

func (csi *ContentServiceImpl) Configure(repository interfaces.ContentRepository) {
	csi.contentRepository = repository
}

func New(contentRepository interfaces.ContentRepository) interfaces.ContentService {
	return &ContentServiceImpl{
		contentRepository: contentRepository,
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
	id, err := csi.contentRepository.Create(post)

	if errors.Is(err, customErrors.ResourceNotCreatedError) || err != nil {
		return 0, customErrors.InternalServerError
	}

	return id, nil

}
