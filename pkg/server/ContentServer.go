package server

import (
	"contentservice/pkg/interfaces/server"
	"contentservice/pkg/server/core/factory"
	"contentservice/pkg/server/log"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ContentServer struct {
	Router *gin.Engine
}

func (server *ContentServer) StartServer(port uint16) error {
	// Init the logger
	log.InitLogger()

	// Start the server
	err := server.Router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	return nil
}

// TODO: Use the "Server" interface
func (server *ContentServer) GetRouter() *gin.Engine {
	return server.Router
}

func NewServer() server.Server {
	return serverFactory.GetContentServer()
}
