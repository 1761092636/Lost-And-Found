package db

import "fmt"

func GetAllUser() (err error, user []User) {
	var sqlStr = `SELECT user_name,address,name,phone FROM tab_user  `
	//获取数据
	rows, err := Db.Query(sqlStr)
	//遍历数据
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.UserName, &u.Address, &u.Name, &u.Phone); err != nil {
			fmt.Println(err)

		}
		user = append(user, u)

	}
	fmt.Println(user)
	if err != nil {

		return err, nil
	}
	return nil, user
}
