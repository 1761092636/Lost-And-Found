package db

// UpdateApplication  更新数据库中应用程序的布尔类型字段
func UpdateApplication(isConfirmed, rewardPaid, result bool, appId int) error {
	// 准备更新语句
	query := "UPDATE tab_applications SET confirmed = ?, reward_paid = ?, result = ? WHERE id = ?"

	// 执行更新语句
	_, err := Db.Exec(query, isConfirmed, rewardPaid, result, appId)
	if err != nil {
		return err
	}

	return nil
}
func UpdateApplication1(isConfirmed bool, appId int) error {
	// 准备更新语句
	query := "UPDATE tab_applications SET confirmed = ?  WHERE id = ?"

	// 执行更新语句
	_, err := Db.Exec(query, isConfirmed, appId)
	if err != nil {
		return err
	}

	return nil
}
