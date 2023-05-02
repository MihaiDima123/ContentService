package main

import (
	int2 "contentservice/application"
)

func main() {
	app := int2.DefaultApp()
	app.Configure()

	app.Start()
}
