package serverInit

import (
	"contentservice/pkg/interfaces/server"
	"contentservice/pkg/serverInit/core/factory"
)

func New() server.Server {
	return serverFactory.GetContentServer()
}
