package handler

import (
	"Lost_and_Found/db"
	"Lost_and_Found/eth"
	"Lost_and_Found/middlewares"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func GetUserPoints(c *gin.Context) {
	todo := db.User{}
	if err := c.ShouldBind(&todo); err != nil {
		middlewares.RespError(c, err)
		return
	}
	arr, err := eth.GetUserPoints(common.HexToAddress(todo.Address))
	if err != nil {
		fmt.Println(err)
		middlewares.RespError(c, err)
	}

	fmt.Println(arr)
	middlewares.RespOK(c, arr)

}
