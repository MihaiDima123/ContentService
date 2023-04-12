package routes

import (
	"contentservice/datasource"
	"github.com/gin-gonic/gin"
)

type Route interface {
	Bind(ginApp *gin.Engine, ds *datasource.Datasource)
}
