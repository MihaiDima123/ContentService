package main

import (
	"contentservice/pkg/application"
	"contentservice/pkg/server/log"
)

func main() {
	app := application.NewDefaultApp().Configure()
	err := app.Start()
	if err != nil {
		log.Error("Could not start the server " + err.Error())
	}
}
