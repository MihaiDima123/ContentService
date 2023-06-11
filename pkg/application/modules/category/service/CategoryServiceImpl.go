package categoryService

import (
	"contentservice/pkg/application/core/customErrors"
	dbErrors "contentservice/pkg/application/core/customErrors/db-errors"
	"contentservice/pkg/application/core/customErrors/httpErrors"
	categoryDto "contentservice/pkg/application/dto/category"
	"contentservice/pkg/application/entity/post_entities"
	"contentservice/pkg/application/modules/category/interfaces"
	"contentservice/pkg/interfaces/customerrors"
)

type CategoryServiceImpl struct {
	repository interfaces.CategoryRepository
}

func (c *CategoryServiceImpl) GetById(id int64) (*categoryDto.CategoryDTO, customerrors.HTTPError) {
	category, err := c.repository.GetById(id)

	if customErrors.Is(err, dbErrors.DbResourceNotFoundType) {
		return nil, httpErrors.HttpNotFoundError
	}

	if err != nil {
		return nil, httpErrors.HttpInternalServerError
	}

	return categoryDto.New(category), nil
}

func (c *CategoryServiceImpl) Create(category *post_entities.Category) (*int64, customerrors.HTTPError) {
	validateError := c.validateCategoryName(category.Name, category.TenantId)

	if validateError != nil {
		return nil, validateError
	}

	id, err := c.repository.Create(category)

	if customErrors.Is(err, dbErrors.DbResourceNotCreatedType) || err != nil {
		return nil, httpErrors.GetInternalServerError(err.Error())
	}

	return id, nil
}

// Validates if the category has a unique name
func (c *CategoryServiceImpl) validateCategoryName(name string, tenantId int) customerrors.HTTPError {
	count, countingError := c.repository.GetCountByNameAndTenant(name, tenantId)

	if countingError != nil {
		return httpErrors.GetInternalServerError(countingError.Error())
	}

	if count != 0 {
		return httpErrors.GetBadRequestError("Category name not unique ", name)
	}

	return nil
}

func New(repository interfaces.CategoryRepository) interfaces.CategoryService {
	service := new(CategoryServiceImpl)
	service.repository = repository
	return service
}
