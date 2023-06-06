package categoryController

import (
	"contentservice/pkg/application/core/customErrors"
	"contentservice/pkg/application/core/customErrors/httpErrors"
	"contentservice/pkg/application/core/parse-param"
	"contentservice/pkg/application/modules/category/interfaces"
	"contentservice/pkg/server/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryController struct {
	service interfaces.CategoryService
}

func (cc *CategoryController) GetById(ctx *gin.Context) {
	id, err := parseParam.GetIntegerParam(parseParam.IdParamName, ctx)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	category, getErr := cc.service.GetById(id)

	if customErrors.Is(getErr, httpErrors.NotFoundErrorType) {
		log.Warn(fmt.Sprintf("Category not found for id %d", id))
		ctx.Status(http.StatusNotFound)
		return
	}

	// TODO: Si create
	ctx.JSON(http.StatusOK, category)
}

func (cc *CategoryController) Create(ctx *gin.Context) {

}

func New(service interfaces.CategoryService) *CategoryController {
	cc := new(CategoryController)
	cc.service = service
	return cc
}
