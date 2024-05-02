package handler

import (
	"Lost_and_Found/eth"
	"github.com/gin-gonic/gin"
)

func Rank(c *gin.Context) {
	// 调用 Rank() 函数获取排名信息列表
	rankList := eth.Rank()

	// 返回 JSON 格式的排名信息列表
	c.JSON(200, rankList)
}
