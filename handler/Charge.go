package handler

import (
	"Lost_and_Found/db"
	"Lost_and_Found/eth"
	"Lost_and_Found/middlewares"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func Charge(c *gin.Context) {

	todo := eth.Tran{}
	if err := c.ShouldBind(&todo); err != nil {
		middlewares.RespError(c, err)
		return
	}

	//根据session当前登陆账户进行操作请求
	var err error
	session := sessions.Default(c)
	var username string
	sessionval := session.Get("username")
	username = sessionval.(string)
	if sessionval == nil {
		// 处理会话值为nil的情况，比如返回错误或执行其他操作
		fmt.Println("Session value is nil")
		middlewares.RespError(c, errors.New("session value is nil"))
		return
	}
	username, ok := sessionval.(string)
	if !ok {
		// 处理类型断言失败的情况，比如返回错误或执行其他操作
		fmt.Println("Session value is not a string")
		middlewares.RespError(c, errors.New("session value is not a string"))
		return
	}
	fmt.Println(username)
	ans, err := db.GetUTCfile(username)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ans)
	reader, _ := os.Open(ans)
	MyLost, Err := bind.NewTransactorWithChainID(reader, "123456", eth.ChainID)
	if Err != nil {
		log.Fatal("NewTransactor", Err)
	}
	MyLost.GasLimit = 3000000
	MyLost.GasPrice = eth.GasPrice
	MyLost.Value = todo.Amount
	fmt.Println("opts:", MyLost)

	_, err = eth.Ins.SendToCaller(MyLost, todo.Amount)
	if err != nil {
		middlewares.RespError(c, err)
		return
	}
	middlewares.RespOK(c, "ok")

}
