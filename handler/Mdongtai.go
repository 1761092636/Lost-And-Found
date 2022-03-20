package handler

import (
	"CDC/CDC-/db"
	"github.com/gin-gonic/gin"
)

func Mdongtai(c *gin.Context) {
	// 解析参数
	todo := db.Med{}
	if err := c.ShouldBind(&todo); err != nil {
		RespError(c, err)
		return
	}

	err, d := db.Dongtai1()
	if err != nil {
		RespError(c, err)
		return
	}

	RespOK(c, d)

}
