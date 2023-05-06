package datasource

import "github.com/jackc/pgx"

type Datasource interface {
	GetConnection() *pgx.ConnPool
	Initialize(configuration Configuration) Datasource
}

type Configuration struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
}
