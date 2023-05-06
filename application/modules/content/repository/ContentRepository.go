package repository

import (
	"contentservice/server/restful/core"
	"github.com/jackc/pgx"
)

type ContentRepository struct {
	dbConn pgx.ConnPool
}

func NewContentRepository() ContentRepository {
	return ContentRepository{}
}

func (cr *ContentRepository) Configure(configuration core.RepositoryConfiguration) {
	cr.dbConn = configuration.Connection
}
