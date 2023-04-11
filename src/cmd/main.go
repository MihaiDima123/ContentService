package main

import (
	"contentservice/initialisation"
)

func init() {
	initialisation.InitEnv()
	initialisation.InitServer()
	initialisation.InitRoutes()
}

func main() {
	initialisation.StartServer()
}
