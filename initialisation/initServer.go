package initialisation

import (
	"github.com/gin-gonic/gin"
)

var router = gin.New()

func InitServer() *gin.Engine {
	return router
}
