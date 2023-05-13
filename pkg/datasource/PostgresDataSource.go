package datasource

import (
	"contentservice/pkg/interfaces/ds"
	"fmt"
	"github.com/jackc/pgx"
)

type PostgresDataSource struct {
	configuration ds.Configuration
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
		// TODO: Panic it's not good, handle your feelings (add a logger)
		panic(err)
	}

	ds.Conn = conn
	fmt.Println(ds.Conn)
}

func (ds *PostgresDataSource) Initialize(configuration ds.Configuration) ds.Datasource {
	appDatasource := &PostgresDataSource{
		configuration: configuration,
	}

	appDatasource.Configure()
	return appDatasource
}

func NewPostgresDataSource() ds.Datasource {
	return new(PostgresDataSource)
}
