package impl

import (
	"contentservice/controller"
	"contentservice/routes"
	"github.com/gin-gonic/gin"
)

type HelloWorldRoute struct {
	Path string
}

func (hwr *HelloWorldRoute) Bind(e *gin.Engine) {
	helloWorldController := controller.HelloWorldController{}

	e.GET(hwr.Path+"/", helloWorldController.GetHelloWorld)
}

func NewHelloWorldRouter(path string) routes.Route {
	helloWorldRoute := &HelloWorldRoute{
		Path: path,
	}
	return helloWorldRoute
}
