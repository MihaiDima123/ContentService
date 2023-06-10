package main

import (
	"contentservice/pkg/application"
	"contentservice/pkg/server/log"
)

func main() {
	app := application.NewDefaultApp().Configure()

	go startUi(app)

	err := app.Start()
	if err != nil {
		log.Error("Could not start the server " + err.Error())
	}
}

func startUi(app *application.App) {
	err := app.StartUiServer()
	if err != nil {
		log.Error("Could not start the UI server " + err.Error())
	}
}
