package main

import (
	"contentservice/pkg/application"
)

func main() {
	app := application.NewDefaultApp().Configure()
	app.Start()
}
