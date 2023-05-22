package content

import (
	"contentservice/pkg/application/modules/content/repository"
	"contentservice/pkg/application/modules/content/services"
	"contentservice/pkg/application/restful/content"
	"contentservice/pkg/interfaces/ds"
	"contentservice/pkg/interfaces/restful"
)

var relativePath = "content"

type CntModule struct {
	Controller content.RestfulContentController
}

func (c *CntModule) Use(configuration *restful.ModuleConfiguration) {
	c.init(configuration.Datasource)

	configuration.Router.GET(relativePath+"/:id", c.Controller.GetPostById)
	configuration.Router.POST(relativePath+"/", c.Controller.CreatePost)
}

func (c *CntModule) init(datasource ds.Datasource) {
	// Repo
	contentRepository := repository.NewContentRepository()
	contentRepository.Configure(restful.RepositoryConfiguration{
		Connection: datasource.GetConnection(),
	})

	// Service
	contentService := services.New(contentRepository)

	contentController := content.NewContentController(contentService)

	c.Controller = contentController
}

func NewContentModule() restful.Module {
	return new(CntModule)
}
