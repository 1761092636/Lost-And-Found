package handler

import (
	"Lost_and_Found/db"
	"Lost_and_Found/middlewares"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
)

func GetUser(c *gin.Context) {

	// 解析参数
	todo := db.User{}
	if err := c.ShouldBind(&todo); err != nil {
		middlewares.RespError(c, err)
		return
	}
	session := sessions.Default(c)
	var username string
	sessionval := session.Get("username")
	username = sessionval.(string)

	fmt.Println(username)

	err, u := db.GetUser(username)
	if err != nil {
		fmt.Println(err)
		return
	}
	//获取余额
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}
	address := u[0].Address
	account := common.HexToAddress(address)
	blockBum, _ := client.BlockNumber(context.Background())
	balance, err := client.BalanceAt(context.Background(), account, big.NewInt(int64(int(blockBum))))
	if err != nil {
		log.Fatal(err)
	}
	u[0].Balance = balance
	middlewares.RespOK(c, u)
}
