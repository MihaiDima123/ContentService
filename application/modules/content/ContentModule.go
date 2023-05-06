package content

import (
	"contentservice/application/modules/content/repository"
	"contentservice/application/modules/content/services/impl"
	"contentservice/datasource"
	"contentservice/server/restful/content"
	serverCore "contentservice/server/restful/core"
)

type CntModule struct {
	Controller content.RestfulContentController
}

func (c *CntModule) Use(configuration *serverCore.ModuleConfiguration) {
	c.init(configuration.Datasource)

	configuration.Router.GET("/", c.Controller.GetHelloWorld)
}

func (c *CntModule) init(datasource datasource.Datasource) {
	contentRepository := repository.NewContentRepository()
	contentRepository.Configure(serverCore.RepositoryConfiguration{
		Connection: *datasource.GetConnection(),
	})

	contentService := impl.NewContentService(contentRepository)

	contentController := content.NewContentController(contentService)

	c.Controller = contentController
}

func NewContentModule() serverCore.Module {
	return new(CntModule)
}
