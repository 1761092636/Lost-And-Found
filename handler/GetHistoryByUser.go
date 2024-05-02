package handler

import (
	"Lost_and_Found/eth"
	"Lost_and_Found/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetHistoryByUser(c *gin.Context) {
	todo := eth.History{}

	if err := c.ShouldBind(&todo); err != nil {
		middlewares.RespError(c, err)
		return
	}
	arr, err := eth.GetHistoryByUser(todo.Operator)
	if err != nil {
		fmt.Println(err)
		middlewares.RespError(c, err)
	}

	fmt.Println(arr)
	middlewares.RespOK(c, arr)

}
