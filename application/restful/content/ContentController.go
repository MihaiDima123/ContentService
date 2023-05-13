package content

import (
	"contentservice/application/modules/content/services"
	"github.com/gin-gonic/gin"
)

type RestfulContentController struct {
	contentService services.ContentService
}

func NewContentController(service services.ContentService) RestfulContentController {
	return RestfulContentController{
		contentService: service,
	}
}

func (rcc *RestfulContentController) GetHelloWorld(context *gin.Context) {
	rcc.contentService.Test()
	_, err := context.Writer.Write([]byte("Hi"))
	if err != nil {
		panic(err)
	}
}
