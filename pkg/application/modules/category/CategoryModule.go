package categoryModule

import (
	"contentservice/pkg/application/modules/category/repository"
	"contentservice/pkg/application/modules/category/service"
	"contentservice/pkg/application/restful/category"
	"contentservice/pkg/interfaces/ds"
	"contentservice/pkg/interfaces/restful"
	"github.com/gin-gonic/gin"
)

type CatMod struct {
	controller *categoryController.CategoryController
}

func (c *CatMod) Use(configuration *restful.ModuleConfiguration) {
	c.initModule(configuration.Datasource)
	c.initRoutes(configuration.Router)
}

func (c *CatMod) initRoutes(router *gin.Engine) {
	router.GET("category/:id", c.controller.GetById)
	router.POST("category", c.controller.Create)
}

func (c *CatMod) initModule(ds ds.Datasource) {
	// Repo
	repo := categoryRepo.New()
	repo.Configure(restful.RepositoryConfiguration{Connection: ds.GetConnection()})

	// Service
	service := categoryService.New(repo)

	// controller
	c.controller = categoryController.New(service)
}

func New() restful.Module {
	return new(CatMod)
}
