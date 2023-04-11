package initialisation

import (
	"contentservice/routes/impl"
	"github.com/gin-gonic/gin"
)

func InitRoutes(e *gin.Engine) {
	impl.NewHelloWorldRouter("/hello-world").Bind(e)
}
