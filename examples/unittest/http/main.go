package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/server/egin"
)

// export EGO_DEBUG=true && go run main.go --config=config.toml
// curl -i 'http://localhost:9006/hello?q=query' -X POST -H 'X-Ego-Uid: 9999' --data '{"id":1,"name":"lee"}'
func main() {
	if err := ego.New().Serve(func() *egin.Component {
		server := egin.Load("server.http").Build()
		server.GET("/hello", GetHelloEgo)
		server.POST("/postHello", PostHelloEgo)
		return server
	}()).Run(); err != nil {
		elog.Panic("startup", elog.FieldErr(err))
	}
}

func GetHelloEgo(c *gin.Context) {
	c.String(200, "Hello EGO")
	return
}

func PostHelloEgo(c *gin.Context) {
	var user struct {
		Name string `json:"name"`
	}
	if err := c.BindJSON(&user); err != nil {
		// bind json出现解析错误后，无法更改状态码
		c.String(http.StatusBadRequest, "invalid params")
		return
	}

	c.String(200, "Hello "+user.Name)
	return
}
