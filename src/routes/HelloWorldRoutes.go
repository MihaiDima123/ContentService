package routes

import (
	"Rome/src/Interface"
	"Rome/src/controller"
	"github.com/gin-gonic/gin"
)

type HelloWorldRoute struct {
	Path string
}

func (hwr *HelloWorldRoute) Bind(e *gin.Engine) {
	helloWorldController := controller.HelloWorldController{}

	e.GET(hwr.Path+"/", helloWorldController.GetHelloWorld)
}

func NewHelloWorldRouter(path string) Interface.Route {
	helloWorldRoute := &HelloWorldRoute{
		Path: path,
	}
	return helloWorldRoute
}
