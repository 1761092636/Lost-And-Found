package handler

import (
	"Lost_and_Found/db"
	"Lost_and_Found/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetUserByAddress(c *gin.Context) {
	todo := db.User{}
	if err := c.ShouldBind(&todo); err != nil {
		middlewares.RespError(c, err)
		return
	}
	arr, err := db.GetUserByAddress(todo.Address)
	if err != nil {
		fmt.Println(err)
		middlewares.RespError(c, err)
	}

	fmt.Println(arr)
	middlewares.RespOK(c, arr)

}
