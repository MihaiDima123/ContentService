package impl

import (
	"contentservice/application/modules/content/repository"
)

type ContentServiceImpl struct {
	contentRepository repository.ContentRepository
}

func NewContentService(contentRepository repository.ContentRepository) ContentServiceImpl {
	return ContentServiceImpl{
		contentRepository: contentRepository,
	}
}
