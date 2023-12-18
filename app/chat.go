package app

import (
	"chat-go/app/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Chat(c *gin.Context) {
	fmt.Println(c.Params)
	// 如何调用刚才封装的 ResponseJson
	util.ResponseOk(c, "hello world")
	// c.JSON(200, gin.H{})
}
