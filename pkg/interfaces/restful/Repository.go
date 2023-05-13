package restful

import (
	"github.com/jackc/pgx"
)

type Repository interface {
	// Configure More like oof, coming from java
	Configure(configuration RepositoryConfiguration)
}

type RepositoryConfiguration struct {
	Connection pgx.ConnPool
}
