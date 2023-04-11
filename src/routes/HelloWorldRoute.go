package routes

import (
	"contentservice/controller"
	"github.com/gin-gonic/gin"
)

type HelloWorldRoute struct {
	Path string
}

func (hwr *HelloWorldRoute) Bind(e *gin.Engine) {
	helloWorldController := controller.HelloWorldController{}

	e.GET(hwr.Path+"/", helloWorldController.GetHelloWorld)
}

func NewHelloWorldRouter(path string) Route {
	helloWorldRoute := &HelloWorldRoute{
		Path: path,
	}
	return helloWorldRoute
}
