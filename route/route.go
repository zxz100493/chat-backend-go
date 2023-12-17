package route

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func NewRoute() *gin.Engine {
	// 强制日志颜色化
	gin.ForceConsoleColor()

	gin.SetMode(gin.ReleaseMode)

	server := gin.Default()
	// server.Use(Cors())
	// server.Use(Recovery)
	group := server.Group("test")
	{
		group.GET("/user", func(ctx *gin.Context) {
			fmt.Println(ctx.Params)
		})
		// group.GET("/user/:uuid", v1.GetUserDetails)
	}
	return server
}
