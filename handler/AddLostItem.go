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
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

// AddLostItem  添加药品
func AddLostItem(c *gin.Context) {
	todo := eth.LostItem{}

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

	err = eth.AddLostItem(MyLost, todo.LostItemName, todo.Description, todo.LostDate, todo.Location, todo.ItemCategory, todo.Reward)
	if err != nil {
		// 处理错误
		fmt.Println("添加丢失物品失败:", err)
		return
	}
	flag := true

	dblost := db.LostItem{}
	err, u := db.GetUser(username)
	if err != nil {
		return
	}
	if err := c.ShouldBind(&dblost); err != nil {
		//middlewares.RespError(c, err)

	}

	// 获取文件并保存到服务器
	file, err := c.FormFile("lostimg")
	if err != nil {
		middlewares.RespError(c, err)
		return
	}
	fileExt := strings.ToLower(path.Ext(file.Filename))
	if fileExt != ".jpg" {
		c.JSON(http.StatusOK, gin.H{"uploading": "done", "message": "上传图片不是jpg格式文件"})
		return
	}
	dst := path.Join("/upload", file.Filename)
	if err := c.SaveUploadedFile(file, "./static/upload/"+file.Filename); err != nil {
		middlewares.RespError(c, "图片上传失败")
		fmt.Println(err)
		return
	}
	dblost.LostImg = dst
	if len(dblost.LostItemName) == 0 {
		middlewares.RespError(c, "失物名称不能为空")
		return
	}

	// 在数据库中插入一条失物数据
	if err := db.AddLostItem(u[0].Address, dblost.LostItemName, dblost.Description, dblost.Location, dblost.ItemCategory, dblost.LostImg, dblost.LostDate, dblost.Reward); err != nil {
		middlewares.RespError(c, err)
		return
	}
	relate, _ := eth.Ins.NextLostItemID(nil)
	relate = big.NewInt(0).Sub(relate, big.NewInt(1))
	fmt.Println(relate)
	if flag {
		err := db.AddHistory(int(relate.Int64()), u[0].Address, "AddlostItem", int(time.Now().Unix()))
		if err != nil {
			fmt.Println(err)
			middlewares.RespError(c, err)

			return
		}
		middlewares.RespOK(c, "添加成功")
	}

}
