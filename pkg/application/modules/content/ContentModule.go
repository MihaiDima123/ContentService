package content

import (
	repoImpl "contentservice/pkg/application/modules/content/repository"
	"contentservice/pkg/application/modules/content/services"
	"contentservice/pkg/application/restful/content"
	"contentservice/pkg/interfaces/ds"
	"contentservice/pkg/interfaces/restful"
)

type CntModule struct {
	Controller content.RestfulContentController
}

func (c *CntModule) Use(configuration *restful.ModuleConfiguration) {
	c.init(configuration.Datasource)

	configuration.Router.GET("/", c.Controller.GetHelloWorld)
}

func (c *CntModule) init(datasource ds.Datasource) {
	contentRepository := repoImpl.NewContentRepository()
	contentRepository.Configure(restful.RepositoryConfiguration{
		Connection: *datasource.GetConnection(),
	})

	contentService := services.NewContentService(contentRepository)

	contentController := content.NewContentController(contentService)

	c.Controller = contentController
}

func NewContentModule() restful.Module {
	return new(CntModule)
}
