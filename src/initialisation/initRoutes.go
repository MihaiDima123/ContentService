package initialisation

import (
	"contentservice/routes/routes"
	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	initRoutes(router)
}

func initRoutes(e *gin.Engine) {
	routes.NewHelloWorldRouter("/hello-world").Bind(e)
}
