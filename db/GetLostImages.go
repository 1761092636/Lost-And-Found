package db

func GetLostImages(key int) (err error, img string) {
	var sqlStr = `SELECT LostImg FROM tab_lostitems WHERE LostItemID =?`
	// 执行查询
	rows, err := Db.Query(sqlStr, key)
	if err != nil {
		// 如果查询出错，返回错误
		return err, ""
	}
	defer rows.Close() // 确保关闭结果集

	// 假设每个LostItemId只对应一个Lostimg，所以我们只读取一行
	if rows.Next() {
		// 读取Lostimg字段的值
		var image string
		err = rows.Scan(&image)
		if err != nil {
			// 如果读取数据出错，返回错误
			println("err")
			return err, ""
		}
		// 返回找到的图像字符串
		return nil, image
	}

	// 如果没有找到匹配的记录，返回没有错误但空字符串
	return nil, ""
}
