package db

// 函数用于根据地址查询失物信息
func GetLostItemsByAddress(address string) ([]LostItem, error) {

	// 查询失物数据
	rows, err := Db.Query("SELECT * FROM tab_lostItems WHERE User = ?", address)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 遍历查询结果并组装成LostItem对象
	var lostItems []LostItem
	for rows.Next() {
		var lostItem LostItem
		err := rows.Scan(
			&lostItem.LostItemID,
			&lostItem.User,
			&lostItem.FoundOne,
			&lostItem.LostItemName,
			&lostItem.Description,
			&lostItem.LostDate,
			&lostItem.Location,
			&lostItem.ItemCategory,
			&lostItem.IsFound,
			&lostItem.FoundDate,
			&lostItem.Reward,
			&lostItem.LostImg,
			&lostItem.FoundApplicationsCount,
		)
		if err != nil {
			return nil, err
		}
		lostItems = append(lostItems, lostItem)
	}

	return lostItems, nil
}
