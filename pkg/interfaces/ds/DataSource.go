package ds

import (
	"database/sql"
)

type Datasource interface {
	GetConnection() *sql.DB
	Initialize(configuration Configuration) Datasource
}

type Configuration struct {
	Host     string
	Port     uint16
	Database string
	User     string
	Password string
}
