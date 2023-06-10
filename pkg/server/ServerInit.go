package server

import (
	"contentservice/pkg/interfaces/server"
	"contentservice/pkg/server/core/factory"
)

func New() server.Server {
	return serverFactory.GetContentServer()
}

func NewUiServer() server.Server {
	return serverFactory.GetUiServer()
}
