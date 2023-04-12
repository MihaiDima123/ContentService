package initialisation

import (
	"contentservice/datasource"
)

func InitDatasource() *datasource.Datasource {
	appDatasource := datasource.NewPostgresDataSource(
		variables.DatasourceHost,
		variables.DatasourcePort,
		variables.DatasourceDatabase,
		variables.DatasourceUser,
		variables.DatasourcePassword,
	)

	appDatasource.Configure()
	return &appDatasource
}
