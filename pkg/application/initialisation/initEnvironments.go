// TODO: Refactor this
package initialisation

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

const appPortName string = "APP_PORT"
const uiPortName string = "UI_PORT"
const datasourceHost string = "DATASOURCE_HOST"
const datasourcePort string = "DATASOURCE_PORT"
const datasourceDatabase string = "DATASOURCE_DATABASE"
const datasourceUser string = "DATASOURCE_USER"
const datasourcePassword string = "DATASOURCE_PASSWORD"

type Environment struct {
	AppPort            uint16
	UiPort             uint16
	DatasourceHost     string
	DatasourcePort     uint16
	DatasourceDatabase string
	DatasourceUser     string
	DatasourcePassword string
}

func InitEnv() Environment {
	// Loads .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	return getEnvs()
}

// Create the object
func getEnvs() Environment {
	return Environment{
		AppPort:            parseIntEnv(os.Getenv(appPortName)),
		UiPort:             parseIntEnv(os.Getenv(uiPortName)),
		DatasourceHost:     os.Getenv(datasourceHost),
		DatasourcePort:     parseIntEnv(os.Getenv(datasourcePort)),
		DatasourceDatabase: os.Getenv(datasourceDatabase),
		DatasourceUser:     os.Getenv(datasourceUser),
		DatasourcePassword: os.Getenv(datasourcePassword),
	}
}

func parseIntEnv(envValue string) uint16 {
	parsedVariable, err := strconv.ParseInt(envValue, 10, 0)
	if err != nil {
		panic(err)
	}
	return uint16(parsedVariable)
}
