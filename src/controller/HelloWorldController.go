package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type dummy struct {
	Test string `json:"test"`
}

type HelloWorldController struct {
}

func (hwc *HelloWorldController) GetHelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, dummy{
		Test: "This is a test",
	})
}
