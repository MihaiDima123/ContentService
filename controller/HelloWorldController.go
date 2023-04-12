package controller

import (
	"contentservice/datasource"
	"github.com/gin-gonic/gin"
)

type HelloWorldController struct {
	Ds datasource.Datasource
}

type Test struct {
	Test string
}

func (hwc *HelloWorldController) GetHelloWorld(c *gin.Context) {
	_, err := hwc.Ds.GetConnection().Exec("insert into test values ('Heloooooo')")
	if err != nil {
		return
	}
}
