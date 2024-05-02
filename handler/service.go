package handler

import (
	"Lost_and_Found/middlewares"
	"github.com/gin-gonic/gin"
)

func Start(addr, webDir string) (err error) {
	// 使用gin框架提供的默认web服务引擎
	r := gin.Default()
	//启用session
	r.Use(middlewares.EnableCookieSession())
	r.Use(middlewares.Cors())
	//获取当前用户

	// 静态文件服务
	if len(webDir) > 0 {
		// 将一个目录下的静态文件，并注册到web服务器

		//r.Static("/", webDir)

	}

	//api接口服务，定义了路由组
	app := r.Group("router")
	{

		////数据库操作
		app.GET("/GetLostItemByIdDb", GetLostItemByIdDb)
		app.GET("/GetLostItemByAddress", GetLostItemByAddress)
		app.POST("/DeleteApplications", DeleteApplications)
		app.POST("/DeleteLostItem", DeleteLostItem)
		app.POST("/UpdateLostItem", UpdateLostItem)
		app.POST("/Login", login)
		app.POST("/LoginOut", LoginOut)
		app.POST("/Register", Register)
		app.GET("/GetUser", GetUser)
		app.POST("/GetAllUser", GetAllUser)
		app.GET("/GetUserByAddress", GetUserByAddress)
		app.POST("/GetApplicationsByReceiver", GetApplicationsByReceiver)
		app.POST("/GetLostItemImages", GetLostItemImages)
		app.POST("/SelectLostItem", SelectLostItem)

		//rank
		app.POST("/Rank", Rank)

		//websocket
		app.GET("/ws", serveWebsocket)

		//eth操作
		app.POST("/GetManger", GetManger)
		app.POST("/GetUserPoints", GetUserPoints)
		app.POST("/GetLostItem", GetLostItem)
		app.GET("/GetLostItemById", GetLostItemById)
		app.POST("/GetHistory", GetHistory)
		app.POST("/GetHistoryByUser", GetHistoryByUser)
		app.POST("/GetApplicationsByReceiverETH", GetApplicationsByReceiverETH)
		app.POST("/Track", Track)
		app.POST("/ConfirmApplication", ConfirmApplication)
		app.POST("/SubmitApplication", SubmitApplication)
		app.POST("/AddLostItem", AddLostItem)
		app.POST("/Charge", Charge)

	}

	// 启动web服务
	err = r.Run(addr)
	return err
}
