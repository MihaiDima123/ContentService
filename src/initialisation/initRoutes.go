package initialisation

import (
	"contentservice/routes"
	"github.com/gin-gonic/gin"
)

func InitRoutes(e *gin.Engine) {
	routes.NewHelloWorldRouter("/hello-world").Bind(e)
}
