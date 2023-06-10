package server

import (
	"github.com/gin-gonic/gin"
)

type Server interface {
	StartServer(port uint16) error
	GetRouter() *gin.Engine
	AddMiddleware(middleware gin.HandlerFunc)
}
