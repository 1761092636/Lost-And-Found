package handler

import (
	"CDC/CDC-/db"
	"github.com/gin-gonic/gin"
)

func Ndongtai(c *gin.Context) {
	// 解析参数
	todo := db.News{}
	if err := c.ShouldBind(&todo); err != nil {
		RespError(c, err)
		return
	}

	err, d := db.Dongtai()
	if err != nil {
		RespError(c, err)
		return
	}

	RespOK(c, d)

}
