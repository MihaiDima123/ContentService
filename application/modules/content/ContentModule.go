package content

import (
	repoImpl "contentservice/application/modules/content/repository/impl"
	"contentservice/application/modules/content/services/impl"
	"contentservice/application/modules/core"
	"contentservice/datasource"
	"contentservice/server/restful/content"
)

type CntModule struct {
	Controller content.RestfulContentController
}

func (c *CntModule) Use(configuration *core.ModuleConfiguration) {
	c.init(configuration.Datasource)

	configuration.Router.GET("/", c.Controller.GetHelloWorld)
}

func (c *CntModule) init(datasource datasource.Datasource) {
	contentRepository := repoImpl.NewContentRepository()
	contentRepository.Configure(core.RepositoryConfiguration{
		Connection: *datasource.GetConnection(),
	})

	contentService := impl.NewContentService(contentRepository)

	contentController := content.NewContentController(contentService)

	c.Controller = contentController
}

func NewContentModule() core.Module {
	return new(CntModule)
}
