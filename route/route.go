package route

import (
	"chat-go/app"

	"github.com/gin-gonic/gin"
)

func NewRoute() *gin.Engine {
	// 强制日志颜色化
	gin.ForceConsoleColor()

	gin.SetMode(gin.ReleaseMode)

	server := gin.Default()
	// server.Use(Cors())
	// server.Use(Recovery)
	group := server.Group("chat")
	{
		group.GET("/test", app.Chat)
		// group.GET("/user/:uuid", v1.GetUserDetails)
	}
	return server
}
