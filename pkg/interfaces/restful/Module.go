package restful

import (
	"contentservice/pkg/interfaces/ds"
	"github.com/gin-gonic/gin"
)

type Module interface {
	Use(configuration *ModuleConfiguration)
}

type ModuleConfiguration struct {
	Datasource ds.Datasource
	Router     gin.Engine
}
