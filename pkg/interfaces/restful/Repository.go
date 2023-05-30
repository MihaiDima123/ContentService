package restful

import (
	"contentservice/pkg/interfaces/errors"
	"gorm.io/gorm"
)

type Repository[T any] interface {
	Configure(configuration RepositoryConfiguration)
	Create(data T) (int64, errors.DbError)
	GetById(id int64) (*T, errors.DbError)
}

type RepositoryConfiguration struct {
	Connection *gorm.DB
}
