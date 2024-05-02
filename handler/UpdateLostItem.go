package handler

import (
	"Lost_and_Found/db"
	"Lost_and_Found/eth"
	"Lost_and_Found/middlewares"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"math/big"
	"os"
)

func UpdateLostItem(c *gin.Context) {
	todo := db.LostItem{}
	if err := c.ShouldBind(&todo); err != nil {
		fmt.Println(err)
		middlewares.RespError(c, err)
		return
	}
	err := db.UpdateLostItem(todo.LostItemID, todo.LostDate, todo.Reward, todo.LostItemName, todo.Description, todo.Location, todo.ItemCategory)
	if err != nil {
		middlewares.RespError(c, err)
		return
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

	err = eth.UpdateLostItem(MyLost, big.NewInt(int64(todo.LostItemID)), todo.LostItemName, todo.Description, big.NewInt(int64(todo.LostDate)), todo.Location, todo.ItemCategory, big.NewInt(int64(todo.Reward)))
	if err != nil {
		// 处理错误
		fmt.Println("更新丢失物品失败:", err)
		middlewares.RespError(c, err)
		return
	}
	middlewares.RespOK(c, "更新成功")
}
