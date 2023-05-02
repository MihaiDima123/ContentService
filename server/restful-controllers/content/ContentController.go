package content

import (
	restControllerInterface "contentservice/server/restful-controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx"
)

func (rcc *RestfulContentController) Configure(options restControllerInterface.Options) {
	rcc.Conn = options.Connection
}

type RestfulContentController struct {
	Conn pgx.ConnPool
}

func (rcc *RestfulContentController) GetHelloWorld(context *gin.Context) {
	fmt.Println(rcc.Conn)

	//_, err := rcc.Conn.Exec("Insert into test (title) values ('test')")
	//if err != nil {
	//	panic(err)
	//}
	//
	//if err != nil {
	//	panic(err)
	//}

	_, err := context.Writer.Write([]byte("Hi"))
	if err != nil {
		panic(err)
	}
}
