package content

import (
	"contentservice/pkg/application/customErrors"
	"contentservice/pkg/application/modules/content/interfaces"
	"contentservice/pkg/server/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RestfulContentController struct {
	contentService interfaces.ContentService
}

func NewContentController(service interfaces.ContentService) RestfulContentController {
	return RestfulContentController{
		contentService: service,
	}
}

func (rcc *RestfulContentController) GetPostById(context *gin.Context) {
	idParam, err := strconv.ParseInt(context.Param("id"), 10, 32)

	if err != nil {
		context.Status(http.StatusNotFound)
		return
	}

	post, err := rcc.contentService.GetById(idParam)

	if httpErr, ok := err.(customErrors.HTTPError); ok {
		log.Warn(fmt.Sprintf("Post not found: %d", idParam))
		context.Status(httpErr.GetStatus())
		return
	}

	if err != nil {
		log.Error(fmt.Sprintf("Error getting the post: %d", idParam))
		context.Status(http.StatusInternalServerError)
		return
	}

	context.JSON(http.StatusOK, post)
}

func (rcc *RestfulContentController) CreatePost(context *gin.Context) {

}
