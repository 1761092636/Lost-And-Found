package db

// Login 登陆
func Login(UserName string, UserPassword string) (err error, a int, b int) {

	var sqlStr = `SELECT user_password FROM tab_user WHERE user_name = ? `
	var sqlStr2 = `SELECT user_name FROM tab_user WHERE user_name = ? `
	//查询id
	var sqlStr1 = `SELECT user_id FROM tab_user WHERE user_name = ?`
	var ps string
	var uname string
	//返回id
	err = Db.QueryRow(sqlStr1, UserName).Scan(&b)
	if err != nil {
		return err, 1, 0
	}

	//绑定值给uname做判断
	err = Db.QueryRow(sqlStr2, UserName).Scan(&uname)
	if err != nil {
		return err, 1, 0
	}
	if uname != UserName {
		return err, 1, 0
	}
	//绑定值给ps做判断
	err = Db.QueryRow(sqlStr, UserName).Scan(&ps)
	if err != nil {
		return err, 1, 0
	}
	if ps == UserPassword {
		return nil, 0, b
	} else if ps != UserPassword {
		return err, 1, 0
	}
	return
}
