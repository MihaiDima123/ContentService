package server

import (
	"contentservice/pkg/interfaces/server"
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

func NewContentServer() server.Server {
	s := new(ContentServer)
	s.Router = gin.Default()

	return s
}
