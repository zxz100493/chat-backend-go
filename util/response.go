package util

import "github.com/gin-gonic/gin"

func responseJson(code int, msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}

func ResponseOk(c *gin.Context, data interface{}, msg ...string) {
	defaultMsg := "Default Message"
	if len(msg) > 0 {
		c.JSON(200, responseJson(0, msg[0], data))
	} else {
		c.JSON(200, responseJson(0, defaultMsg, data))
	}
}

func ResponseError(c *gin.Context, data interface{}, msg ...string) {
	defaultMsg := "Default Error"
	if len(msg) > 0 {
		c.JSON(200, responseJson(-1, msg[0], data))
	} else {
		c.JSON(200, responseJson(-1, defaultMsg, data))
	}
}
