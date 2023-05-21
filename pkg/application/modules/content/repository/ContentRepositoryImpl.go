package repository

import (
	"contentservice/pkg/application/modules/content/interfaces"
	"contentservice/pkg/interfaces/restful"
	"gorm.io/gorm"
)

type ContentRepositoryImpl struct {
	dbConn *gorm.DB
}

func (cr *ContentRepositoryImpl) Create() {
}

func (cr *ContentRepositoryImpl) GetById(id int) {

}

func NewContentRepository() interfaces.ContentRepository {
	return new(ContentRepositoryImpl)
}

func (cr *ContentRepositoryImpl) Configure(configuration restful.RepositoryConfiguration) {
	cr.dbConn = configuration.Connection
}
