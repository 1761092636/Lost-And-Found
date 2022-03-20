package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init() (err error) {
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/cdc")
	if err != nil {
		return err
	} else {
		fmt.Println("连接数据库成功")
	}

	return
}

// Dongtai 动态展示新闻
func Dongtai() (err error, news []News) {
	var sqlStr = `SELECT * FROM tab_news `
	//获取数据
	rows, err := db.Query(sqlStr)
	//遍历数据
	for rows.Next() {
		var n News
		if err := rows.Scan(&n.NewsId, &n.Content, &n.Title, &n.NewsImg, &n.Author, &n.Timeset, &n.Like); err != nil {
			fmt.Println(err)

		}
		news = append(news, n)

	}
	fmt.Println(news)

	if err != nil {

		return err, nil
	}
	return nil, news
}

// Dongtai1 动态展示药品
func Dongtai1() (err error, med []Med) {
	var sqlStr = `SELECT * FROM tab_med`
	//获取数据
	rows, err := db.Query(sqlStr)
	//遍历数据
	for rows.Next() {
		var m Med
		if err := rows.Scan(&m.MedId, &m.MedName, &m.MedTxt, &m.MedPrice, &m.MedImg); err != nil {
			fmt.Println(err)

		}
		med = append(med, m)

	}
	fmt.Println(med)

	if err != nil {

		return err, nil
	}
	return nil, med

}
