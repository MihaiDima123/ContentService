package initialisation

import (
	"contentservice/datasource"
	"contentservice/routes"
	"github.com/gin-gonic/gin"
)

func InitRoutes(ds *datasource.Datasource) {
	initRoutes(router, ds)
}

func initRoutes(e *gin.Engine, ds *datasource.Datasource) {
	routes.NewHelloWorldRouter("/hello-world").Bind(e, ds)
}
