package serverFactory

import (
	serverInterface "contentservice/pkg/interfaces/server"
	"contentservice/pkg/server/core/servers"
	"github.com/gin-gonic/gin"
)

func GetContentServer() serverInterface.Server {
	cs := new(servers.ContentServer)
	cs.Router = gin.Default()

	return cs
}
