package handler

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("username")
	err := session.Save()
	if err != nil {
		return
	}

}
