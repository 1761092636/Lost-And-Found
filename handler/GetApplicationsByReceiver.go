package handler

import (
	"Lost_and_Found/db"
	"Lost_and_Found/middlewares"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func GetApplicationsByReceiver(c *gin.Context) {
	todo := db.Application{}
	if err := c.ShouldBind(&todo); err != nil {
		middlewares.RespError(c, err)
		return
	}
	arr, err := db.GetApplicationsByReceiver(todo.Receiver)
	if err != nil {
		fmt.Println(err)
		middlewares.RespError(c, err)
		return
	}

	session := sessions.Default(c)
	var username string
	sessionval := session.Get("username")
	if sessionval != nil {
		username = sessionval.(string)
	} else {
		middlewares.RespError(c, fmt.Errorf("username not found in session"))
		return
	}

	// 遍历每一个申请，如果isconfirm字段为false，则发送WebSocket消息
	for _, app := range arr {
		if !app.IsConfirmed { // 假设Isconfirm是db.Application结构体的字段
			// 将int类型的时间戳转换为time.Time
			submitTime := time.Unix(int64(app.SubmitDate), 0)

			// 格式化时间为更易读的字符串形式
			const layout = "2006年1月2日 15:04" // Go的布局字符串，用于指定日期和时间的格式
			readableTime := submitTime.Format(layout)

			item := "#" + strconv.Itoa(app.RelateItemID)

			// 构造WebSocket消息
			message := "new application message on " + readableTime + " on " + item

			// 发送WebSocket消息
			if err := SendMessageToUser(username, message); err != nil {
				fmt.Println("Error sending WebSocket message:", err)
				// 你可以根据错误类型决定是否返回错误给客户端
			}
		}
	}

	// 发送成功，返回申请列表给客户端，无论是否发送消息
	middlewares.RespOK(c, arr)
}
