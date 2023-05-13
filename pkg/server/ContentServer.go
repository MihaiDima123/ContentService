package server

import (
	"contentservice/pkg/server/log"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ContentServer struct {
	Router *gin.Engine
}

func (server *ContentServer) StartServer(port uint16) {
	// Init the logger
	log.InitLogger()

	// Start the server
	err := server.Router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		return
	}
}

func (server *ContentServer) GetRouter() *gin.Engine {
	return server.Router
}

func NewServer() *ContentServer {
	return GetContentServer()
}
