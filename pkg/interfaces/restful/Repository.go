package restful

import (
	"contentservice/pkg/application/core/customErrors"
	"gorm.io/gorm"
)

type Repository[T any] interface {
	Configure(configuration RepositoryConfiguration)
	Create(data T) (int64, customErrors.DbError)
	GetById(id int64) (*T, customErrors.DbError)
}

type RepositoryConfiguration struct {
	Connection *gorm.DB
}
