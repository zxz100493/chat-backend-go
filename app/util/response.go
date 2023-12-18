package util

import "github.com/gin-gonic/gin"

// 封装一个gin框架的json返回方法
func responseJson(code int, msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}

// 字节来一个gin框架的返回 让我可以不用再写c.JSON(200 这种格式
// 可以把这个c参数也省略吗

func ResponseOk(c *gin.Context, data interface{}, msg ...string) {
	defaultMsg := "Default Message"
	if len(msg) > 0 {
		c.JSON(200, responseJson(0, msg[0], data))
	} else {
		c.JSON(200, responseJson(0, defaultMsg, data))
	}
}
