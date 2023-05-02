package server

import (
	"contentservice/helpers"
	"github.com/gin-gonic/gin"
)

type ContentServer struct {
	Port   int64
	Router *gin.Engine
}

func (server *ContentServer) ConfigureServer(port int64) {
	server.Port = port
	server.Router = gin.Default()
}

func (server *ContentServer) StartServer() {
	err := server.Router.Run(helpers.GetPort(server.Port))
	if err != nil {
		return
	}
}

func (server *ContentServer) GetRouter() *gin.Engine {
	return server.Router
}

func NewContentServer() Server {
	return &ContentServer{}
}
