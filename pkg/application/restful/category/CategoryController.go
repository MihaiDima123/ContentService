package categoryController

import (
	"contentservice/pkg/application/core/customErrors"
	"contentservice/pkg/application/core/customErrors/httpErrors"
	validationErrors "contentservice/pkg/application/core/customErrors/validation-errors"
	idWrapper "contentservice/pkg/application/core/entities/wrappers"
	"contentservice/pkg/application/core/parse-param"
	problemDetailImpl "contentservice/pkg/application/core/problemDetail"
	validationInstance "contentservice/pkg/application/core/validator/instance"
	categoryDto "contentservice/pkg/application/dto/category"
	"contentservice/pkg/application/modules/category/interfaces"
	"contentservice/pkg/interfaces/customerrors"
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
	category, bindErr := getCategoryFromBody(ctx)

	if bindErr != nil {
		log.Error(bindErr.Error())
		ctx.JSON(http.StatusBadRequest, problemDetailImpl.NewOfValidationError(bindErr))
		return
	}

	id, err := cc.service.Create(category)
	if err != nil {
		log.OfStatus(err.GetStatus(), err.Error())
		ctx.JSON(err.GetStatus(), problemDetailImpl.NewOfHttpError(err).Title("Could not create category"))
		return
	}
	ctx.JSON(http.StatusCreated, idWrapper.Of(id))
	return
}

func getCategoryFromBody(ctx *gin.Context) (*categoryDto.CategoryDTO, customerrors.ValidationError) {
	category := new(categoryDto.CategoryDTO)

	bindErr := ctx.BindJSON(&category)
	if bindErr != nil {
		return nil, validationErrors.GetParseError(validationInstance.New("Body", bindErr.Error()))
	}
	return category, nil
}

func New(service interfaces.CategoryService) *CategoryController {
	cc := new(CategoryController)
	cc.service = service
	return cc
}
