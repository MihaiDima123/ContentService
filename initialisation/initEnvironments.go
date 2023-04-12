package initialisation

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

const appPortName string = "APP_PORT"
const datasourceHost string = "DATASOURCE_HOST"
const datasourcePort string = "DATASOURCE_PORT"
const datasourceDatabase string = "DATASOURCE_DATABASE"
const datasourceUser string = "DATASOURCE_USER"
const datasourcePassword string = "DATASOURCE_PASSWORD"

type environment struct {
	AppPort            int64
	DatasourceHost     string
	DatasourcePort     uint16
	DatasourceDatabase string
	DatasourceUser     string
	DatasourcePassword string
}

var variables *environment

func InitEnv() {
	// Loads .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	variables = getEnvs()
}

// Create the object
func getEnvs() *environment {
	envs := environment{
		AppPort:            parseIntEnv(os.Getenv(appPortName)),
		DatasourceHost:     os.Getenv(datasourceHost),
		DatasourcePort:     uint16(parseIntEnv(os.Getenv(datasourcePort))),
		DatasourceDatabase: os.Getenv(datasourceDatabase),
		DatasourceUser:     os.Getenv(datasourceUser),
		DatasourcePassword: os.Getenv(datasourcePassword),
	}

	return &envs
}

func parseIntEnv(envValue string) int64 {
	parsedVariable, err := strconv.ParseInt(envValue, 10, 0)
	if err != nil {
		panic(err)
	}
	return parsedVariable
}
