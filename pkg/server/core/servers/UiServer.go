package servers

import (
	"contentservice/pkg/server/log"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UiServer struct {
	Router *gin.Engine
}

func (server *UiServer) StartServer(port uint16) error {
	// Init the logger
	log.InitLogger()

	server.Router.Static("/", "./static")

	// Start the server
	err := server.Router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	return nil
}

func (server *UiServer) GetRouter() *gin.Engine {
	return server.Router
}

func (server *UiServer) AddMiddleware(middleware gin.HandlerFunc) {
	server.Router.Use(middleware)
}
