package core

import (
	"contentservice/datasource"
	"github.com/gin-gonic/gin"
)

type Module interface {
	Use(configuration *ModuleConfiguration)
}

type ModuleConfiguration struct {
	Datasource datasource.Datasource
	Router     gin.Engine
}
