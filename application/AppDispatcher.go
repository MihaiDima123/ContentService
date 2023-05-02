package application

type Dispatcher struct {
}

func (ad *Dispatcher) DispatchRoutes(app *App) {
	contentModule := NewContentModule()

	app.use(&contentModule)
}
