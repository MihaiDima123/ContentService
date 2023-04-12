package datasource

import "github.com/jackc/pgx"

type Datasource interface {
	Configure()
	GetConnection() *pgx.Conn
}
