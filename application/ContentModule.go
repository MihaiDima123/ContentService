package application

import (
	"contentservice/server/restful-controllers"
	"contentservice/server/restful-controllers/content"
	"github.com/gin-gonic/gin"
)

type TypeContentModule struct {
	controller content.RestfulContentController
}

func (tcm *TypeContentModule) Bind(app *App) {
	tcm.controller.Configure(restful_controllers.Options{
		Connection: *app.DataSource.GetConnection(),
	})

	app.Server.GetRouter().GET("/", func(context *gin.Context) {
		tcm.controller.GetHelloWorld(context)
	})
}

func GetController() content.RestfulContentController {
	return content.RestfulContentController{}
}

func NewContentModule() TypeContentModule {
	return TypeContentModule{
		controller: GetController(),
	}
}
