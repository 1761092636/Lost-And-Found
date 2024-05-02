package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

// Init 初始化数据库
func Init() (err error) {
	Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/lostandfound")
	if err != nil {
		return err
	} else {
		fmt.Println("连接数据库成功")
	}

	return
}

//// Dongtai 动态展示新闻
//func Dongtai() (err error, news []News) {
//	var sqlStr = `SELECT * FROM tab_news `
//	//获取数据
//	rows, err := db.Query(sqlStr)
//	//遍历数据
//	for rows.Next() {
//		var n News
//		if err := rows.Scan(&n.NewsId, &n.Content, &n.Title, &n.NewsImg, &n.Author, &n.Timeset); err != nil {
//			fmt.Println(err)
//
//		}
//		news = append(news, n)
//
//	}
//	fmt.Println(news)
//
//	if err != nil {
//
//		return err, nil
//	}
//	return nil, news
//}
//
//// Dongtai1 动态展示药品
//func Dongtai1() (err error, med []Med) {
//	var sqlStr = `SELECT * FROM tab_med`
//	//获取数据
//	rows, err := db.Query(sqlStr)
//	//遍历数据
//	for rows.Next() {
//		var m Med
//		if err := rows.Scan(&m.MedId, &m.MedName, &m.MedTxt, &m.MedPrice, &m.MedImg); err != nil {
//			fmt.Println(err)
//
//		}
//		med = append(med, m)
//
//	}
//	fmt.Println(med)
//
//	if err != nil {
//
//		return err, nil
//	}
//	return nil, med
//
//}
//
//// Dongtai2 动态购物车商品
//func Dongtai2() (err error, med []Med) {
//	var sqlStr = `select   b.* from tab_cart as a
//left join tab_med  b  on  a.p_id = b.med_id`
//	//获取数据
//	rows, err := db.Query(sqlStr)
//	//遍历数据
//	for rows.Next() {
//		var m Med
//		if err := rows.Scan(&m.MedId, &m.MedName, &m.MedTxt, &m.MedPrice, &m.MedImg); err != nil {
//			fmt.Println(err)
//
//		}
//		med = append(med, m)
//
//	}
//	fmt.Println(med)
//
//	if err != nil {
//
//		return err, nil
//	}
//	return nil, med
//}
//
//// Dongtai3 动态展示科普文章
//func Dongtai3() (err error, blog []Blog) {
//	var sqlStr = `SELECT * FROM tab_blog `
//	//获取数据
//	rows, err := db.Query(sqlStr)
//	//遍历数据
//	for rows.Next() {
//		var b Blog
//		if err := rows.Scan(&b.BlogId, &b.BlogTitle, &b.BlogContent, &b.BlogImg); err != nil {
//			fmt.Println(err)
//
//		}
//		blog = append(blog, b)
//
//	}
//	if err != nil {
//
//		return err, nil
//	}
//	return nil, blog
//
//}
//
//// SearchBlog 查找科普
//func SearchBlog(key string) (error, []Blog) {
//	var sqlStr = `SELECT * FROM tab_blog WHERE  blog_title LIKE '%` + key + `%'  `
//	//获取数据
//	rows, err := db.Query(sqlStr)
//	if err != nil {
//
//		return err, nil
//	}
//	//遍历数据
//	var blog []Blog
//	for rows.Next() {
//		var b Blog
//		if err := rows.Scan(&b.BlogId, &b.BlogTitle, &b.BlogContent, &b.BlogImg); err != nil {
//			fmt.Println(err)
//		}
//		blog = append(blog, b)
//
//	}
//	return nil, blog
//
//}
//
//// SearchNews 查找新闻
//func SearchNews(key string) (error, []News) {
//	var sqlStr = `SELECT * FROM tab_news WHERE   title LIKE '%` + key + `%' OR  author LIKE '%` + key + `%'`
//	//获取数据
//	rows, err := db.Query(sqlStr)
//	if err != nil {
//
//		return err, nil
//	}
//	//遍历数据
//	var news []News
//	for rows.Next() {
//		var n News
//		if err := rows.Scan(&n.NewsId, &n.Title, &n.Content, &n.NewsImg, &n.Author, &n.Timeset); err != nil {
//			fmt.Println(err)
//		}
//		news = append(news, n)
//
//	}
//	return nil, news
//
//}
//

//// AddBlog 发布科普
//func AddBlog(Blog_title string, Blog_content string, Blog_img string) (err error) {
//	var sqlStr = `INSERT INTO tab_blog (blog_title, blog_content, blog_img) VALUES (?,?,?) `
//	_, err = db.Exec(sqlStr, Blog_title, Blog_content, Blog_img)
//	if err != nil {
//		return err
//	}
//
//	return
//}
//
//// DeleteNews 删除新闻
//func DeleteNews(NewsId int) (err error) {
//	var sqlStr = `DELETE FROM tab_news WHERE news_id = ?`
//
//	_, err = db.Exec(sqlStr, NewsId)
//	if err != nil {
//		return err
//	}
//
//	return
//}
//
//// DeleteBlog 删除科普
//func DeleteBlog(BlogId int) (err error) {
//	var sqlStr = `DELETE FROM tab_blog WHERE blog_id = ?`
//
//	_, err = db.Exec(sqlStr, BlogId)
//	if err != nil {
//		return err
//	}
//
//	return
//
//}
//
//// EditNews 修改新闻
//
//func GetEditNews(key int) (error, []News) {
//	var sqlStr = `select * from tab_news where news_id= ? `
//
//	rows, err := db.Query(sqlStr, key)
//	if err != nil {
//		return err, nil
//	}
//	var news []News
//	for rows.Next() {
//		var n News
//		if err := rows.Scan(&n.NewsId, &n.Title, &n.Content, &n.NewsImg, &n.Author, &n.Timeset); err != nil {
//			fmt.Println(err)
//		}
//		news = append(news, n)
//
//	}
//	return nil, news
//
//}
//
//func EditNews(Content string, Title string, NewsImg string, Author string, Timeset string, NewsId int) (err error) {
//	var sqlStr = `UPDATE tab_news SET  content = ? ,title = ? ,news_img =? ,author = ?,timeset = ? WHERE news_id = ?`
//
//	_, err = db.Exec(sqlStr, Content, Title, NewsImg, Author, Timeset, NewsId)
//	if err != nil {
//		return err
//	}
//
//	return
//}
//
//// EditBlog 修改科普
//
//func GetEditBlog(key int) (error, []Blog) {
//	var sqlStr = `SELECT * FROM tab_blog WHERE blog_id = ?`
//	//获取数据
//	rows, err := db.Query(sqlStr, key)
//	if err != nil {
//
//		return err, nil
//	}
//	//遍历数据
//	var blog []Blog
//	for rows.Next() {
//		var b Blog
//		if err := rows.Scan(&b.BlogId, &b.BlogTitle, &b.BlogContent, &b.BlogImg); err != nil {
//			fmt.Println(err)
//		}
//		blog = append(blog, b)
//
//	}
//	return nil, blog
//
//}
//
//func EditBlog(Title string, Content string, BlogImg string, BlogId int) (err error) {
//	var sqlStr = `UPDATE tab_blog SET  blog_title = ? ,blog_content = ? ,blog_img = ?  WHERE blog_id = ?`
//
//	_, err = db.Exec(sqlStr, Title, Content, BlogImg, BlogId)
//	if err != nil {
//		return err
//
//	}
//	if Title == "" {
//		return err
//	}
//	if Content == "" {
//		return err
//
//	}
//
//	return
//}
//
//// GetAdvice 获取意见反馈
//func GetAdvice() (err error, adv []Adv) {
//	var sqlStr = `SELECT * FROM tab_adv `
//	//获取数据
//	rows, err := db.Query(sqlStr)
//	//遍历数据
//	for rows.Next() {
//		var a Adv
//		if err := rows.Scan(&a.AdvId, &a.Name, &a.Address, &a.Phone, &a.DrugName, &a.Hash, &a.Info); err != nil {
//			fmt.Println(err)
//
//		}
//		adv = append(adv, a)
//
//	}
//	fmt.Println(adv)
//
//	if err != nil {
//
//		return err, nil
//	}
//	return nil, adv
//}
//
//// Advice 增加意见反馈
//func Advice(Name string, Address string, Phone int, DrugName string, Hash string, Info string) (err error) {
//	var sqlStr = `INSERT INTO tab_adv (your_name, your_address, your_phone,your_drug_name,your_hash,info) VALUES (?,?,?,?,?,?) `
//	_, err = db.Exec(sqlStr, Name, Address, Phone, DrugName, Hash, Info)
//	if err != nil {
//		return err
//	}
//
//	return
//
//}
//
//func SearchAdvice(key string) (error, []Adv) {
//	var sqlStr = `SELECT * FROM tab_adv WHERE  your_name LIKE '%` + key + `%' OR  your_drug_name LIKE '%` + key + `%'`
//	//获取数据
//	rows, err := db.Query(sqlStr)
//	if err != nil {
//
//		return err, nil
//	}
//	//遍历数据
//	var adv []Adv
//	for rows.Next() {
//		var a Adv
//		if err := rows.Scan(&a.AdvId, &a.Name, &a.Address, &a.Phone, &a.DrugName, &a.Hash, &a.Info); err != nil {
//			fmt.Println(err)
//		}
//		adv = append(adv, a)
//
//	}
//	return nil, adv
//
//}
