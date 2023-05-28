package restful

import (
	"contentservice/pkg/application/core/customErrors/interfaces"
	"gorm.io/gorm"
)

type Repository[T any] interface {
	Configure(configuration RepositoryConfiguration)
	Create(data T) (int64, interfaces.DbError)
	GetById(id int64) (*T, interfaces.DbError)
}

type RepositoryConfiguration struct {
	Connection *gorm.DB
}
