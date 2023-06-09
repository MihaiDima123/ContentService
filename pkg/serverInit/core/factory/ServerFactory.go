package serverFactory

import (
	serverInterface "contentservice/pkg/interfaces/server"
	"contentservice/pkg/serverInit/core/servers"
	"github.com/gin-gonic/gin"
)

func GetContentServer() serverInterface.Server {
	cs := new(servers.ContentServer)
	cs.Router = gin.Default()

	return cs
}
