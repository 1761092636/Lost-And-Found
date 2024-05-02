package handler

import (
	"Lost_and_Found/db"
	"Lost_and_Found/middlewares"
	"github.com/gin-gonic/gin"
)

func Chat(c *gin.Context) {
	todo := db.Message{}
	if err := c.ShouldBind(&todo); err != nil {
		middlewares.RespError(c, err)
		return
	}

}
