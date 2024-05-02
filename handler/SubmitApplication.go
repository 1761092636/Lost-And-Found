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
	"os"
	"strconv"
	"time"
)

func SubmitApplication(c *gin.Context) {
	todo := eth.Application{}
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

	err = eth.SubmitApplication(MyLost, todo.RelateItemID)
	if err != nil {
		// 处理错误
		fmt.Println("申请失败:", err)
		return
	} else {
		owner, err := eth.Ins.GetLostItemOwner(nil, todo.RelateItemID)
		if err != nil {
			return
		}
		touser, _ := db.GetUserByAddress(owner.String())
		const layout = "2006年1月2日 15:04" // Go的布局字符串，用于指定日期和时间的格式
		readableTime := time.Now().Format(layout)
		itme, _ := db.GetLostItemByID(int(todo.RelateItemID.Int64()))
		message := "REALTIME: new application message on" + readableTime + "on #" + strconv.Itoa(itme.LostItemID) + itme.LostItemName
		err = SendMessageToUser(touser.UserName, message)
		if err != nil {
			return
		}
		middlewares.RespOK(c, nil)
	}

}
