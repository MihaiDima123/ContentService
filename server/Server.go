package server

import (
	"github.com/gin-gonic/gin"
)

type Server interface {
	ConfigureServer(port int64)
	StartServer()
	GetRouter() *gin.Engine
}
