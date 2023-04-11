package main

import (
	"Rome/src/initialisation"
	"Rome/src/initialisation/env"
)

func init() {
	env.InitEnv()
	initialisation.InitServer()
}

func main() {
	initialisation.StartServer()
}
