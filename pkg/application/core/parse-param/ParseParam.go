package parseParam

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

const (
	IdParamName = "id"
)

// GetParam Not implemented anywhere yet...
func GetParam(name string, ctx *gin.Context) string {
	return ctx.Param(name)
}

func GetIntegerParam(name string, ctx *gin.Context) (int64, error) {
	stringParam := ctx.Param(name)
	return parseStringToInteger(stringParam)
}

// Some utils
func parseStringToInteger(value string) (int64, error) {
	return strconv.ParseInt(value, 10, 32)
}
