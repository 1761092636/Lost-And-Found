package db

import (
	"database/sql"
	"errors"
)

// GetUserByAddress 通过地址查询用户信息
func GetUserByAddress(address string) (User, error) {
	var user User

	// 执行数据库查询操作
	row := Db.QueryRow("SELECT user_id, user_name, address, name, phone FROM tab_user WHERE address = ?", address)

	// 将查询结果填充到User结构体中
	err := row.Scan(&user.UserId, &user.UserName, &user.Address, &user.Name, &user.Phone)
	if err != nil {
		// 如果未找到用户，则返回自定义错误
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, errors.New("user not found")
		}
		// 如果发生其他错误，则直接返回
		return User{}, err
	}

	return user, nil
}
