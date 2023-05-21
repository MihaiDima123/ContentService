package datasource

import (
	"contentservice/pkg/interfaces/ds"
	"contentservice/pkg/server/log"
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type PostgresDataSource struct {
	configuration ds.Configuration
	Conn          *sql.DB
}

func (ds *PostgresDataSource) GetConnection() *sql.DB {
	return ds.Conn
}

func (ds *PostgresDataSource) Configure() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		ds.configuration.Host,
		ds.configuration.User,
		ds.configuration.Password,
		ds.configuration.Database,
		ds.configuration.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	databaseConn, _ := db.DB()
	databaseConn.SetMaxIdleConns(10)
	databaseConn.SetMaxOpenConns(20)
	databaseConn.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Error("Could not open database connection" + err.Error())
		panic(err)
	}

	ds.Conn = databaseConn
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
