package routes

import "github.com/gin-gonic/gin"

type Route interface {
	Bind(ginApp *gin.Engine)
}
