package services

import (
	"contentservice/pkg/application/modules/content/repository"
)

type ContentServiceImpl struct {
	contentRepository repository.ContentRepository
}

func (c ContentServiceImpl) Test() {
	c.contentRepository.Test()
}

func NewContentService(contentRepository repository.ContentRepository) ContentService {
	return ContentServiceImpl{
		contentRepository: contentRepository,
	}
}
