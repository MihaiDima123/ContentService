package datasource

import (
	"github.com/jackc/pgx"
	_ "github.com/jackc/pgx"
)

type PostgresDataSource struct {
	host     string
	port     uint16
	database string
	user     string
	password string
	Conn     *pgx.Conn
}

func (ds *PostgresDataSource) GetConnection() *pgx.Conn {
	return ds.Conn
}

func (ds *PostgresDataSource) Configure() {
	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     ds.host,
		Port:     ds.port,
		Database: ds.database,
		User:     ds.user,
		Password: ds.password,
	})

	if err != nil {
		panic(err)
	}

	ds.Conn = conn
}

func NewPostgresDataSource(
	host string,
	port uint16,
	database string,
	user string,
	password string,
) Datasource {
	return &PostgresDataSource{
		host:     host,
		port:     port,
		database: database,
		user:     user,
		password: password,
	}
}
