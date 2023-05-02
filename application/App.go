package application

import (
	"contentservice/datasource"
	"contentservice/env/initialisation"
	"contentservice/server"
)

type App struct {
	Server       server.Server
	DataSource   datasource.Datasource
	environments *initialisation.Environment
}

func (app *App) Configure() {
	app.environments = initialisation.InitEnv()
	app.DataSource = *app.getDataSource()

	appServer := server.NewContentServer()
	appServer.ConfigureServer(app.environments.AppPort)

	app.Server = appServer
}

func (app *App) Start() {
	appDispatcher := Dispatcher{}
	appDispatcher.DispatchRoutes(app)

	// Finally
	app.Server.StartServer()
}

func (app *App) use(module Module) {
	module.Bind(app)
}

func (app *App) getDataSource() *datasource.Datasource {
	ds := datasource.NewPostgresDataSource(datasource.Configuration{
		Host:     app.environments.DatasourceHost,
		Port:     app.environments.DatasourcePort,
		Database: app.environments.DatasourceDatabase,
		User:     app.environments.DatasourceUser,
		Password: app.environments.DatasourcePassword,
	})
	ds.Initialize()

	return &ds
}

func DefaultApp() App {
	return App{}
}
