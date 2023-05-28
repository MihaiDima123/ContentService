package content

import (
	errorsInterface "contentservice/pkg/application/core/customErrors/interfaces"
	"contentservice/pkg/application/entity/post_entities"
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
		log.Warn("Bad request, could not parse id param")
		context.Status(http.StatusBadRequest)
		return
	}

	post, err := rcc.contentService.GetById(idParam)

	if httpErr, ok := err.(errorsInterface.HTTPError); ok {
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
	post := post_entities.Post{}
	err := context.BindJSON(&post)

	if err != nil {
		log.Warn("error when parsing request body" + err.Error())
		context.Status(http.StatusBadRequest)
		return
	}

	id, err := rcc.contentService.Create(post)

	if httpErr, ok := err.(errorsInterface.HTTPError); ok {
		log.Warn("Could not create post: " + httpErr.Error())
		context.Status(httpErr.GetStatus())
		return
	}

	context.JSON(http.StatusOK, id)
}
