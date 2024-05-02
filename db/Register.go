package db

// Register 注册
func Register(UserName string, UserPassword string, Address string, utcfile string, Name string, Phone string) (err error) {
	var sqlStr = `INSERT INTO tab_user (user_name,user_password,address,utcfile,name,phone) VALUES (?,?,?,?,?,?)`
	_, err = Db.Exec(sqlStr, UserName, UserPassword, Address, utcfile, Name, Phone)
	if err != nil {
		return err
	}

	return
}
