package application

import (
	"contentservice/application/initialisation"
	"contentservice/application/modules/content"
	"contentservice/application/modules/core"
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

	/*
		Then the server - I have just one but just in case this part will allow migrating to a new api version inside this
		project cause why not, maybe a feature switch added a TODO: Feature switch
	*/
	app.Server = server.NewContentServer()
	// Should work on its own please don't add anything more to that class thanks

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
