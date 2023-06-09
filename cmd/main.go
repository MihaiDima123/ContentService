package main

import (
	"contentservice/pkg/application"
	"contentservice/pkg/serverInit/log"
)

func main() {
	app := application.NewDefaultApp().Configure()
	err := app.Start()
	if err != nil {
		log.Error("Could not start the server " + err.Error())
	}
}
