package env

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type environment struct {
	AppPort int64
}

var Variables *environment

func InitEnv() {
	// Loads .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	Variables = getEnvs()
}

// Create the object
func getEnvs() *environment {
	envs := environment{
		AppPort: parseIntEnv(os.Getenv(appPortName)),
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
