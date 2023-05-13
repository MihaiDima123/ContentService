package impl

import (
	"contentservice/application/modules"
	"contentservice/application/modules/content/repository"
	"github.com/jackc/pgx"
)

type ContentRepositoryImpl struct {
	dbConn pgx.ConnPool
}

func (cr *ContentRepositoryImpl) Test() {
	_, err := cr.dbConn.Exec("Insert into test values ('This is a test')")
	if err != nil {
		return
	}
}

func NewContentRepository() repository.ContentRepository {
	return new(ContentRepositoryImpl)
}

func (cr *ContentRepositoryImpl) Configure(configuration modules.RepositoryConfiguration) {
	cr.dbConn = configuration.Connection
}
