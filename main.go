package main

import (
	"CDC/CDC-/config"
	"CDC/CDC-/db"
	"CDC/CDC-/handler"
	"fmt"
	"os"
)

func main() {
	var err error
	//eth.Ect()
	// 载入配置文件

	if len(os.Args) != 2 {
		fmt.Printf("usage: %v config-file\n", os.Args[0])
		return
	}
	c, err := config.LoadConfig(os.Args[1])
	if err != nil {
		fmt.Printf("载入配置文件错误:%v\n", err)
		return
	}
	fmt.Println(*c)

	//初始化数据库
	err = db.Init()
	if err != nil {
		fmt.Printf("初始化数据库错误:%v\n", err)
		return
	}

	//开启web服务器

	err = handler.Start(fmt.Sprintf("%s:%s", c.Host, c.Port), c.WebDir)
	fmt.Println(c.WebDir)
	if err != nil {
		fmt.Printf("web服务错误:%v\n", err)

		return
	}

}
