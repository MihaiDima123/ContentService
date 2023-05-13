package server

import (
	"github.com/gin-gonic/gin"
)

type Server interface {
	StartServer(port uint16)
	GetRouter() *gin.Engine
}
