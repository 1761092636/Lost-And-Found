package db

func GetLostItemByID(id int) (*LostItem, error) {
	// 准备查询语句
	query := "SELECT * FROM tab_lostitems WHERE LostItemID = ?"

	// 执行查询
	row := Db.QueryRow(query, id)

	// 准备变量来存储查询结果
	var item LostItem

	// 扫描查询结果到结构体
	err := row.Scan(
		&item.LostItemID,
		&item.User,
		&item.FoundOne,
		&item.LostItemName,
		&item.Description,
		&item.LostDate,
		&item.Location,
		&item.ItemCategory,
		&item.IsFound,
		&item.FoundDate,
		&item.Reward,
		&item.LostImg,
		&item.FoundApplicationsCount,
	)
	if err != nil {
		return nil, err
	}

	return &item, nil
}
