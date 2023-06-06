package serverFactory

import (
	"contentservice/pkg/server"
	"github.com/gin-gonic/gin"
)

func GetContentServer() *server.ContentServer {
	cs := new(server.ContentServer)
	cs.Router = gin.Default()

	return cs
}
