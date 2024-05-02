package handler

import (
	"Lost_and_Found/eth"
	"Lost_and_Found/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetHistory(c *gin.Context) {
	arr, err := eth.GetHistory()
	if err != nil {
		fmt.Println(err)
		middlewares.RespError(c, err)
	}

	fmt.Println(arr)
	middlewares.RespOK(c, arr)

}
