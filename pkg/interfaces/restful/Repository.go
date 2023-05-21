package restful

import (
	"gorm.io/gorm"
)

type Repository interface {
	Configure(configuration RepositoryConfiguration)
	Create()
	GetById(id int)
}

type RepositoryConfiguration struct {
	Connection *gorm.DB
}
