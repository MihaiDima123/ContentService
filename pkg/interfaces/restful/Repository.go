package restful

import (
	"database/sql"
)

type Repository interface {
	Configure(configuration RepositoryConfiguration)
}

type RepositoryConfiguration struct {
	Connection *sql.DB
}
