package impl

import (
	"contentservice/pkg/application/modules/content/repository"
	"contentservice/pkg/application/modules/content/services"
)

type ContentServiceImpl struct {
	contentRepository repository.ContentRepository
}

func (c ContentServiceImpl) Test() {
	c.contentRepository.Test()
}

func NewContentService(contentRepository repository.ContentRepository) services.ContentService {
	return ContentServiceImpl{
		contentRepository: contentRepository,
	}
}
