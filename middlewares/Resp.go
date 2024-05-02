package middlewares

import "github.com/gin-gonic/gin"

// RespOK 封装函数，用于向客户端返回正确信息
func RespOK(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 0,
		"data": data,
	})
}

// RespError 封装函数，用于向客户端返回错误消息
func RespError(c *gin.Context, msg interface{}) {
	c.JSON(200, gin.H{
		"code":    1,
		"message": msg,
	})
}
