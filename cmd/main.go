package main

import (
	"contentservice/application"
)

func main() {
	app := application.NewDefaultApp().Configure()
	app.Start()
}
