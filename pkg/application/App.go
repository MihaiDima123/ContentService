package application

import (
	"contentservice/pkg/application/initialisation"
	categoryModule "contentservice/pkg/application/modules/category"
	"contentservice/pkg/application/modules/content"
	"contentservice/pkg/datasource"
	"contentservice/pkg/interfaces/ds"
	"contentservice/pkg/interfaces/restful"
	serverInt "contentservice/pkg/interfaces/server"
	"contentservice/pkg/server"
	"contentservice/pkg/server/log"
)

type App struct {
	Server       serverInt.Server
	UiServer     serverInt.Server
	DataSource   ds.Datasource
	environments initialisation.Environment
}

// TODO: Apply a strategy there
func (app *App) Configure() *App {
	app.environments = initialisation.InitEnv() // Init the environments
	app.DataSource = app.getDataSource()        // Then the data source

	// Backend server
	app.Server = server.New()

	// Swag
	app.UiServer = server.NewUiServer()

	// Content
	content.NewContentModule().Use(&restful.ModuleConfiguration{
		Datasource: app.DataSource,
		Router:     app.Server.GetRouter(),
	})

	// Category
	categoryModule.New().Use(&restful.ModuleConfiguration{
		Datasource: app.DataSource,
		Router:     app.Server.GetRouter(),
	})

	return app
}

func (app *App) Start() error {
	err := app.Server.StartServer(app.environments.AppPort)
	if err != nil {
		log.Error("Failed to start the server")
		return err
	}
	return nil
}

func (app *App) StartUiServer() error {
	err := app.UiServer.StartServer(app.environments.UiPort)
	if err != nil {
		log.Error("Failed to start the UI server")
		return err
	}
	return nil
}

func (app *App) getDataSource() ds.Datasource {
	return datasource.NewPostgresDataSource().Initialize(ds.Configuration{
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
