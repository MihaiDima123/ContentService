package categoryRepo

import (
	dbErrors "contentservice/pkg/application/core/customErrors/db-errors"
	"contentservice/pkg/application/entity/post_entities"
	"contentservice/pkg/application/modules/category/interfaces"
	"contentservice/pkg/interfaces/customerrors"
	"contentservice/pkg/interfaces/restful"
	"errors"
	"gorm.io/gorm"
)

const CategoryTableName = "category"

type CategoryRepository struct {
	dbConn *gorm.DB
}

func (c CategoryRepository) Configure(configuration restful.RepositoryConfiguration) {
	c.dbConn = configuration.Connection
}

func (c CategoryRepository) Create(data post_entities.Category) (int64, customerrors.DbError) {
	result := c.dbConn.Table(CategoryTableName).Create(&data)
	if result.Error != nil {
		return 0, dbErrors.DbNotCreatedError
	}

	return data.ID, nil
}

func (c CategoryRepository) GetById(id int64) (*post_entities.Category, customerrors.DbError) {
	category := new(post_entities.Category)
	err := c.dbConn.Table(CategoryTableName).Where("id = ?", id).First(category).Error

	if errors.Is(err, gorm.ErrRecordNotFound) || err != nil {
		return nil, dbErrors.DbNotFoundError
	}

	return category, nil
}

func New() interfaces.CategoryRepository {
	return new(CategoryRepository)
}
