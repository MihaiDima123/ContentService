package helpers

import "fmt"

func GetPort(portValue uint16) string {
	return fmt.Sprintf(":%d", portValue)
}
