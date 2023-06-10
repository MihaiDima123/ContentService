package categoryController

import (
	"contentservice/pkg/application/core/customErrors"
	"contentservice/pkg/application/core/customErrors/httpErrors"
	"contentservice/pkg/application/core/parse-param"
	problemDetailImpl "contentservice/pkg/application/core/problemDetail"
	"contentservice/pkg/application/entity/post_entities"
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

	ctx.JSON(http.StatusOK, category)
}

func (cc *CategoryController) Create(ctx *gin.Context) {
	category := new(post_entities.Category)

	bindError := ctx.BindJSON(&category)
	if bindError != nil {
		log.Error(fmt.Sprintf("%s for category %v", bindError.Error(), category))
		ctx.JSON(http.StatusBadRequest,
			problemDetailImpl.NewOfHttpError(httpErrors.HttpBadRequestError).
				Detail(bindError.Error()))
		return
	}

	id, err := cc.service.Create(category)
	if err != nil {
		log.Warn(fmt.Sprintf("Could not create category %v", err.Error()))
		ctx.JSON(http.StatusInternalServerError,
			problemDetailImpl.NewOfHttpError(httpErrors.HttpInternalServerError).
				Detail(err.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, id)
	return
}

func New(service interfaces.CategoryService) *CategoryController {
	cc := new(CategoryController)
	cc.service = service
	return cc
}
