package handler

import (
	"Lost_and_Found/eth"
	"Lost_and_Found/middlewares"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func GetManger(c *gin.Context) {
	manger := eth.GetManger()
	var empty common.Address
	if manger == empty {
		middlewares.RespError(c, "空地址")
	}
	middlewares.RespOK(c, manger)
}
