package initialisation

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func InitServer() {
	router = gin.Default()
}

func StartServer() {
	err := router.Run(getPort(Variables.AppPort))
	if err != nil {
		return
	}
}

func getPort(portValue int64) string {
	return fmt.Sprintf(":%d", portValue)
}
