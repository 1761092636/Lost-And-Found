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

func DeleteLostItem(c *gin.Context) {
	flag := true
	var err error
	todo := db.LostItem{}
	if err = c.ShouldBind(&todo); err != nil {
		middlewares.RespError(c, err)
		return
	} else {
		err = db.DeleteLostItem(todo.LostItemID)
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
		var lostid = big.NewInt(int64(todo.LostItemID))

		err = eth.DeleteLostItems(MyLost, lostid)
		if err != nil {
			// 处理错误
			flag = false
			fmt.Println("删除丢失物品失败:", err)
			return
		}

	}

	if flag == true {
		middlewares.RespOK(c, "删除成功")
	}

}
