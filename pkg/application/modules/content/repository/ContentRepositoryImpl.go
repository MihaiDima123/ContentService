package repository

import (
	"contentservice/pkg/application/modules/content/interfaces"
	"contentservice/pkg/interfaces/restful"
	"database/sql"
)

type ContentRepositoryImpl struct {
	dbConn *sql.DB
}

func (cr *ContentRepositoryImpl) Test() {
	_, err := cr.dbConn.Exec("Insert into test values ('This is a test')")
	if err != nil {
		return
	}
}

func NewContentRepository() interfaces.ContentRepository {
	return new(ContentRepositoryImpl)
}

func (cr *ContentRepositoryImpl) Configure(configuration restful.RepositoryConfiguration) {
	cr.dbConn = configuration.Connection
}
