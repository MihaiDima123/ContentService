package services

import "contentservice/pkg/application/modules/content/repository"

type TemplateServiceImpl struct {
	repository *repository.ContentRepositoryImpl
}
