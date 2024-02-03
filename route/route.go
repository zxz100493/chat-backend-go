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

	group := server.Group("chat")
	{
		group.GET("/test", app.Chat)
	}

	return server
}
