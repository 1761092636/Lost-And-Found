package db

// AddLostItem 添加失物
func AddLostItem(user, lostItemName, description, location, itemCategory, lostimg string, lostDate, reward int) error {
	query := "INSERT INTO tab_lostItems (User,  LostItemName, Description, LostDate, Location, ItemCategory,Lostimg, Reward) VALUES (?, ?, ?, ?, ?, ?, ?,?)"
	_, err := Db.Exec(query, user, lostItemName, description, lostDate, location, itemCategory, lostimg, reward)
	if err != nil {
		return err
	}

	return nil
}
