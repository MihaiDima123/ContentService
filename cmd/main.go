package main

import (
	"contentservice/initialisation"
)

func init() {
	initialisation.InitEnv()
	initialisation.InitServer()

	// Routes + datasource
	ds := initialisation.InitDatasource()
	initialisation.InitRoutes(ds)
}

func main() {
	initialisation.StartServer()
}
