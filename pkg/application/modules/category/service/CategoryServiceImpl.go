package categoryService

import (
	"contentservice/pkg/application/core/customErrors"
	dbErrors "contentservice/pkg/application/core/customErrors/db-errors"
	"contentservice/pkg/application/core/customErrors/httpErrors"
	"contentservice/pkg/application/entity/post_entities"
	"contentservice/pkg/application/modules/category/interfaces"
	"contentservice/pkg/interfaces/customerrors"
)

type CategoryServiceImpl struct {
	repository interfaces.CategoryRepository
}

func (c *CategoryServiceImpl) GetById(id int64) (*post_entities.Category, customerrors.HTTPError) {
	category, err := c.repository.GetById(id)

	if customErrors.Is(err, dbErrors.DbResourceNotFoundType) {
		return nil, httpErrors.HttpNotFoundError
	}

	if err != nil {
		return nil, httpErrors.HttpInternalServerError
	}

	return category, nil
}

func (c *CategoryServiceImpl) Create(category *post_entities.Category) (int64, customerrors.HTTPError) {
	id, err := c.repository.Create(category)

	if customErrors.Is(err, dbErrors.DbResourceNotCreatedType) || err != nil {
		return 0, httpErrors.GetInternalServerError(err.Error())
	}

	return id, nil
}

func New(repository interfaces.CategoryRepository) interfaces.CategoryService {
	service := new(CategoryServiceImpl)
	service.repository = repository
	return service
}
