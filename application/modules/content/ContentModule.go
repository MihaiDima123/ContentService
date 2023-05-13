package content

import (
	"contentservice/application/modules"
	repoImpl "contentservice/application/modules/content/repository/impl"
	"contentservice/application/modules/content/services/impl"
	"contentservice/application/restful/content"
	"contentservice/datasource"
)

type CntModule struct {
	Controller content.RestfulContentController
}

func (c *CntModule) Use(configuration *modules.ModuleConfiguration) {
	c.init(configuration.Datasource)

	configuration.Router.GET("/", c.Controller.GetHelloWorld)
}

func (c *CntModule) init(datasource datasource.Datasource) {
	contentRepository := repoImpl.NewContentRepository()
	contentRepository.Configure(modules.RepositoryConfiguration{
		Connection: *datasource.GetConnection(),
	})

	contentService := impl.NewContentService(contentRepository)

	contentController := content.NewContentController(contentService)

	c.Controller = contentController
}

func NewContentModule() modules.Module {
	return new(CntModule)
}
