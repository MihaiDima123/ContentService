package routes

import (
	"contentservice/controller"
	"contentservice/datasource"
	"github.com/gin-gonic/gin"
)

type HelloWorldRoute struct {
	Path string
}

func (hwr *HelloWorldRoute) Bind(e *gin.Engine, ds *datasource.Datasource) {
	helloWorldController := controller.HelloWorldController{
		Ds: *ds,
	}

	e.GET(hwr.Path+"/", helloWorldController.GetHelloWorld)
}

func NewHelloWorldRouter(path string) Route {
	helloWorldRoute := &HelloWorldRoute{
		Path: path,
	}
	return helloWorldRoute
}
