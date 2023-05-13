package server

import (
	"github.com/gin-gonic/gin"
)

func GetContentServer() *ContentServer {
	cs := new(ContentServer)
	cs.Router = gin.Default()

	return cs
}
