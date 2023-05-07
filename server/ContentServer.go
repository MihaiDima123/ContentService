package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type ContentServer struct {
	Router *gin.Engine
}

func (server *ContentServer) StartServer(port uint16) {
	err := server.Router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		return
	}
}

func (server *ContentServer) GetRouter() *gin.Engine {
	return server.Router
}

func NewContentServer() Server {
	server := new(ContentServer)
	server.Router = gin.Default()

	return server
}
