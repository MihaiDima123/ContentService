package restful

import (
	"github.com/jackc/pgx"
)

type Repository interface {
	Configure(configuration RepositoryConfiguration)
}

type RepositoryConfiguration struct {
	Connection pgx.ConnPool
}
