package db

// GetUTCfile 获取utc文件存放牡蛎
func GetUTCfile(username string) (a string, err error) {

	var sqlstr = `SELECT utcfile FROM tab_user WHERE user_name = ?`
	err = Db.QueryRow(sqlstr, username).Scan(&a)
	if err != nil {
		return "", err
	}
	return a, nil
}
