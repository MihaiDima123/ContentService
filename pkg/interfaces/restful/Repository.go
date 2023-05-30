package restful

import (
	"contentservice/pkg/interfaces/customerrors"
	"gorm.io/gorm"
)

type Repository[T any] interface {
	Configure(configuration RepositoryConfiguration)
	Create(data T) (int64, customerrors.DbError)
	GetById(id int64) (*T, customerrors.DbError)
}

type RepositoryConfiguration struct {
	Connection *gorm.DB
}
