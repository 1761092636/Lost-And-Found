package handler

import (
	"Lost_and_Found/db"
	"Lost_and_Found/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	err, arr := db.GetAllUser()
	if err != nil {
		fmt.Println(err)
		middlewares.RespError(c, err)
	}
	fmt.Println(arr)
	middlewares.RespOK(c, arr)

}
