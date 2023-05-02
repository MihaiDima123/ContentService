package datasource

import (
	"fmt"
	"github.com/jackc/pgx"
)

type PostgresDataSource struct {
	configuration Configuration
	Conn          *pgx.ConnPool
}

func (ds *PostgresDataSource) GetConnection() *pgx.ConnPool {
	return ds.Conn
}

func (ds *PostgresDataSource) Configure() {
	conn, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     ds.configuration.Host,
			Port:     ds.configuration.Port,
			Database: ds.configuration.Database,
			User:     ds.configuration.User,
			Password: ds.configuration.Password,
		},
		MaxConnections: 10,
		AcquireTimeout: 2000,
	})

	if err != nil {
		panic(err)
	}

	ds.Conn = conn
	fmt.Println(ds.Conn)

}

func (ds *PostgresDataSource) Initialize() Datasource {
	appDatasource := NewPostgresDataSource(ds.configuration)
	appDatasource.Configure()

	return appDatasource
}

func NewPostgresDataSource(configuration Configuration) Datasource {
	pgDataSource := PostgresDataSource{}
	pgDataSource.configuration = configuration

	return &pgDataSource
}
