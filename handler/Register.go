package handler

import (
	"Lost_and_Found/db"
	"Lost_and_Found/middlewares"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	// 解析参数
	todo := db.User{}
	if err := c.ShouldBind(&todo); err != nil {
		middlewares.RespError(c, err)
		return
	}
	//设置报错提示
	if len(todo.UserName) == 0 {
		middlewares.RespError(c, "用户名不能为空")
		return
	}
	if len(todo.UserPassword) == 0 {
		middlewares.RespError(c, "密码不能为空")
		return
	}
	//创建以太坊账户
	//保存keystore文件在指定目录下
	ks := keystore.NewKeyStore("./keytstore", keystore.StandardScryptN, keystore.StandardScryptP)
	newAcc, _ := ks.NewAccount("123456")

	//强制转换newacc中的address为string类型
	MyAcc := fmt.Sprintf("%v", newAcc.Address)
	todo.Address = MyAcc
	MyUtc := fmt.Sprintf("%v", newAcc.URL.Path)
	todo.UtcFile = MyUtc
	// 在数据库中插入用户信息

	if err := db.Register(todo.UserName, todo.UserPassword, todo.Address, todo.UtcFile, todo.Name, todo.Phone); err != nil {
		middlewares.RespError(c, err)
		return
	}

	middlewares.RespOK(c, "添加成功")

}
