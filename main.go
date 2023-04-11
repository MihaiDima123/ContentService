package main

import "Rome/initialisation"

func init() {
	initialisation.InitServer()
}

func main() {
	myServer := GetRouter()

	err := myServer.Run(":8080")
	if err != nil {
		return
	}
}
