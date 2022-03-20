package handler

import (
	"github.com/gin-gonic/gin"
)

func Start(addr, webDir string) (err error) {
	// 使用gin框架提供的默认web服务引擎
	r := gin.Default()
	// 静态文件服务
	if len(webDir) > 0 {
		// 将一个目录下的静态文件，并注册到web服务器

		r.Static("/web", webDir)
	}

	//api接口服务，定义了路由组
	app := r.Group("router")
	{
		app.POST("/news", Ndongtai)
		app.POST("/products", Mdongtai)
		// 	// 定义增改查的接口，并注册到web服务器

	}

	// 启动web服务
	err = r.Run(addr)
	return err
}

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
