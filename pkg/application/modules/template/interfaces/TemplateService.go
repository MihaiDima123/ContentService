package interfaces

import "contentservice/pkg/interfaces/restful"

type TemplateService interface {
	restful.Service
	Configure(repository TemplateRepository)
}
