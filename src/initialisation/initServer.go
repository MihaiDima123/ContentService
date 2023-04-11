package initialisation

import (
	"contentservice/initialisation/env"
	"fmt"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func InitServer() {
	// Create the router
	router = gin.Default()

	// Bind the routes
	InitRoutes(router)
}

func StartServer() {
	err := router.Run(getPort(env.Variables.AppPort))
	if err != nil {
		return
	}
}

func getPort(portValue int64) string {
	return fmt.Sprintf(":%d", portValue)
}
