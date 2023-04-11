package main

import (
	"contentservice/initialisation"
	"contentservice/initialisation/env"
)

func init() {
	env.InitEnv()
	initialisation.InitServer()
}

func main() {
	initialisation.StartServer()
}
