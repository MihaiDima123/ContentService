package main

import (
	"contentservice/initialisation"
)

func init() {
	initialisation.InitEnv()
	initialisation.InitServer()
}

func main() {
	initialisation.StartServer()
}
