package datasource

import (
	"contentservice/pkg/interfaces/ds"
	"contentservice/pkg/server/log"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDataSource struct {
	configuration ds.Configuration
	Conn          *gorm.DB
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

func (ds *PostgresDataSource) GetConnection() *gorm.DB {
	return ds.Conn
}

func (ds *PostgresDataSource) Configure() {

	db, err := gorm.Open(
		postgres.Open(getConnectionString(ds.configuration)), new(gorm.Config))

	if err != nil {
		log.Error("Could not open database connection" + err.Error())
		panic(err)
	}

	databaseConn, dbConnErr := db.DB()
	databaseConn.SetMaxIdleConns(10)
	databaseConn.SetMaxOpenConns(100)

	if dbConnErr != nil {
		log.Error("Database error" + dbConnErr.Error())
		panic(dbConnErr)
	}

	ds.Conn = db
}

func getConnectionString(configuration ds.Configuration) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		configuration.Host,
		configuration.User,
		configuration.Password,
		configuration.Database,
		configuration.Port)
}
