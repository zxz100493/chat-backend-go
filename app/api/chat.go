package api

import (
	"chat-go/services"
	"chat-go/util"

	"github.com/gin-gonic/gin"
)

func Chat(c *gin.Context) {
	msg := c.Query("msg")

	if msg == "" {
		util.ResponseError(c, "no message")
		return
	}

	res := services.ChatWithAi(msg)

	util.ResponseOk(c, res)
}
