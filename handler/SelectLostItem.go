package handler

import (
	"Lost_and_Found/db"
	"Lost_and_Found/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SelectLostItem(c *gin.Context) {
	todo := db.LostItem{}
	if err := c.ShouldBind(&todo); err != nil {
		middlewares.RespError(c, err)
		return
	}
	err, arr := db.SelectLostItem(todo.ItemCategory)
	if err != nil {
		fmt.Println(err)
		middlewares.RespError(c, err)
	}

	fmt.Println(arr)
	middlewares.RespOK(c, arr)

}
