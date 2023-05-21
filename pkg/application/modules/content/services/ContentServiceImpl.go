package services

import (
	"contentservice/pkg/application/modules/content/interfaces"
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
