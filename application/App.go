package application

import (
	"contentservice/application/core"
	"contentservice/application/initialisation"
	"contentservice/application/modules/content"
	"contentservice/datasource"
	"contentservice/server"
)

type App struct {
	Server       server.Server
	DataSource   datasource.Datasource
	environments initialisation.Environment
}

func (app *App) Configure() *App {
	app.environments = initialisation.InitEnv() // Init the environments
	app.DataSource = app.getDataSource()        // Then the data source

	app.Server = server.NewContentServer()

	content.NewContentModule().Use(&core.ModuleConfiguration{
		Datasource: app.DataSource,
		Router:     *app.Server.GetRouter(),
	})

	return app
}

func (app *App) Start() {
	app.Server.StartServer(app.environments.AppPort)
}

func (app *App) getDataSource() datasource.Datasource {
	return datasource.NewPostgresDataSource().Initialize(datasource.Configuration{
		Host:     app.environments.DatasourceHost,
		Port:     app.environments.DatasourcePort,
		Database: app.environments.DatasourceDatabase,
		User:     app.environments.DatasourceUser,
		Password: app.environments.DatasourcePassword,
	})
}

func NewDefaultApp() *App {
	return new(App)
}
