package handler

import (
	"Lost_and_Found/db"
	"Lost_and_Found/eth"
	"Lost_and_Found/middlewares"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	"os"
	"time"
)

func DeleteApplications(c *gin.Context) {
	flag := true
	var err error

	todo := db.Application{}
	if err = c.ShouldBind(&todo); err != nil {
		middlewares.RespError(c, err)
		return
	}
	err = db.DeleteApplication(todo.ApplicationId)
	if err != nil {
		fmt.Println(err)
		flag = false
		middlewares.RespError(c, err)
	}

	//根据session当前登陆账户进行操作请求
	session := sessions.Default(c)
	var username string
	sessionval := session.Get("username")
	username = sessionval.(string)
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
	fmt.Println("opts:", MyLost)

	err = eth.DeleteApplication(MyLost, big.NewInt(int64(todo.RelateItemID)), big.NewInt(int64(todo.ApplicationId)))

	if err != nil {
		// 处理错误
		flag = false
		fmt.Println("删除申请失败:", err)
		middlewares.RespError(c, err)
		return
	}
	if flag {
		err, myuser := db.GetUser(username)
		fmt.Println(common.HexToAddress(myuser[0].Address))
		err = db.AddHistory(todo.RelateItemID, myuser[0].Address, "DeleteApplication", int(time.Now().Unix()))
		if err != nil {
			fmt.Println(err)
			middlewares.RespError(c, err)
			return
		}
		middlewares.RespOK(c, "删除成功")
	}

}
