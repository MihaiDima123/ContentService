package content

import (
	"contentservice/pkg/application/modules/content/interfaces"
	"github.com/gin-gonic/gin"
)

type RestfulContentController struct {
	contentService interfaces.ContentService
}

func NewContentController(service interfaces.ContentService) RestfulContentController {
	return RestfulContentController{
		contentService: service,
	}
}

func (rcc *RestfulContentController) GetHelloWorld(context *gin.Context) {
	_, err := context.Writer.Write([]byte("Hi"))
	if err != nil {
		panic(err)
	}
}
