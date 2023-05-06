package content

import (
	"contentservice/application/modules/content/services"
	"github.com/gin-gonic/gin"
)

type RestfulContentController struct {
	contentService services.ContentService
}

func NewContentController(cs services.ContentService) RestfulContentController {
	return RestfulContentController{
		contentService: cs,
	}
}

func (rcc *RestfulContentController) GetHelloWorld(context *gin.Context) {
	_, err := context.Writer.Write([]byte("Hi"))
	if err != nil {
		panic(err)
	}
}
