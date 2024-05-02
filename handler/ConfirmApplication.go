package handler

import (
	"Lost_and_Found/db"
	"Lost_and_Found/eth"
	"Lost_and_Found/middlewares"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

func ConfirmApplication(c *gin.Context) {
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
	reader, _ := os.Open(ans)
	MyLost, Err := bind.NewTransactorWithChainID(reader, "123456", eth.ChainID)
	if Err != nil {
		log.Fatal("NewTransactor", Err)
	}
	lost, _ := eth.Ins.LostItems(nil, todo.RelateItemID)
	//获取reward值
	MyLost.Value = lost.Reward
	MyLost.GasLimit = 3000000
	MyLost.GasPrice = eth.GasPrice
	fmt.Println("opts:", MyLost)

	traceHash, err := eth.ConfirmApplication(MyLost, todo.RelateItemID, todo.ApplicationId, todo.Result)
	if err != nil {
		// 处理错误
		fmt.Println("操作失败:", err)
		middlewares.RespError(c, err)
		return
	}

	err, myuser := db.GetUser(username)
	if err != nil {
		return
	}
	apps, err := eth.GetApplicationsByReceiver(common.HexToAddress(myuser[0].Address))
	fmt.Println(common.HexToAddress(myuser[0].Address))
	if err != nil {
		fmt.Println(err)

		return
	}

	//操作数据库
	// 遍历接收者收到的所有申请
	for _, app := range apps {
		// 如果当前申请已被确认

		if app.ApplicationId.Cmp(todo.ApplicationId) == 0 && app.RelateItemID.Cmp(todo.RelateItemID) == 0 {
			fmt.Println(app)
			if todo.Result == true {
				// 更新当前申请的状态
				err = db.UpdateApplication(app.IsConfirmed, app.RewardPaid, app.Result, int(app.ApplicationId.Int64()))
				if err != nil {
					middlewares.RespError(c, err)
					return
				}
				// 更新失物信息
				fmt.Println(int(time.Now().Unix()))
				err = db.UpdateFound(app.Sender.String(), int(app.RelateItemID.Int64()), int(time.Now().Unix()))
				if err != nil {
					middlewares.RespError(c, err)
					return
				}
				for _, otherApp := range apps {
					// 排除当前申请本身
					if otherApp.ApplicationId.Cmp(app.ApplicationId) != 0 && otherApp.RelateItemID.Cmp(app.RelateItemID) == 0 {

						err = db.UpdateApplication1(true, int(otherApp.ApplicationId.Int64()))
						if err != nil {
							middlewares.RespError(c, err)
							return
						}

					}
				}
			}
			if todo.Result == false {
				err = db.UpdateApplication1(app.IsConfirmed, int(app.ApplicationId.Int64()))
				if err != nil {
					return
				}
			}
		}
	}

	// 如果调用成功，打印成功消息
	fmt.Println("操作成功")
	err = db.AddHistory(int(todo.RelateItemID.Int64()), myuser[0].Address, "DoneFound", int(time.Now().Unix()))
	if err != nil {
		fmt.Println(err)
		middlewares.RespError(c, err)
		return
	}
	middlewares.RespOK(c, traceHash)
}
