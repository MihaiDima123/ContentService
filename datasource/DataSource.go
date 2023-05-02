package datasource

import "github.com/jackc/pgx"

type Datasource interface {
	Configure()
	GetConnection() *pgx.ConnPool
	Initialize() Datasource
}

type Configuration struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
}
