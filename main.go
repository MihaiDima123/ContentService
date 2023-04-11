package main

import (
	"Rome/initialisation"
	"Rome/initialisation/env"
)

func init() {
	env.InitEnv()
	initialisation.InitServer()
}

func main() {
	initialisation.StartServer()
}
