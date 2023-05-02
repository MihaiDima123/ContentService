package helpers

import "fmt"

func GetPort(portValue int64) string {
	return fmt.Sprintf(":%d", portValue)
}
