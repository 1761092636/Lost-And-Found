package handler

import (
	"Lost_and_Found/db"
	"Lost_and_Found/middlewares"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	session := sessions.Default(c)
	// 解析参数
	todo := db.User{}
	if err := c.ShouldBind(&todo); err != nil {
		middlewares.RespError(c, err)
		return
	}
	// 在数据库中匹配用户信息
	err, a, b := db.Login(todo.UserName, todo.UserPassword)
	if err != nil {
		middlewares.RespError(c, err.Error())
		return
	}
	if a == 0 {

		session.Set("username", todo.UserName)
		println(session.Get("username").(string))
		err := session.Save()
		if err != nil {
			return
		}
		middlewares.RespOK(c, b)
		return
	} else {
		middlewares.RespError(c, "请检查用户名或密码")

	}

}
