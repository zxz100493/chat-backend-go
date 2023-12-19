package app

import (
	"chat-go/app/util"
	"chat-go/infrastructurre/ai/baidu"

	"github.com/gin-gonic/gin"
)

func Chat(c *gin.Context) {
	msg := c.Query("msg")

	if msg == "" {
		util.ResponseError(c, "no message")
		return
	}

	baidu := baidu.Baidu{}
	res := baidu.Chat(msg)
	util.ResponseOk(c, res)
}
