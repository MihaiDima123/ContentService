package application

type Module interface {
	Bind(app *App)
}
