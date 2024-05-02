package db

import (
	"fmt"
)

// UpdateLostItem 更新数据库中指定失物ID的失物信息
func UpdateLostItem(lostItemID, LostDate, Reward int, LostItemName, Description, Location, ItemCategory string) error {
	// 准备SQL语句
	query := "UPDATE tab_lostitems SET LostItemName=?, Description=?, LostDate=?, Location=?, ItemCategory=?,  Reward=? WHERE LostItemID=?"

	// 执行 SQL 语句插入数据
	_, err := Db.Exec(query, LostItemName, Description, LostDate, Location, ItemCategory, Reward, lostItemID)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("插入数据失败: %v", err)
	}

	fmt.Printf("失物ID为 %d 的失物信息已成功更新\n", lostItemID)
	return nil
}
func UpdateFound(foundone string, LostItemID, FoundDate int) error {
	// 准备SQL语句
	query := "UPDATE tab_lostitems SET FoundOne =?, IsFound=?,FoundDate =?  WHERE LostItemID=?"

	// 执行 SQL 语句插入数据
	_, err := Db.Exec(query, foundone, true, FoundDate, LostItemID)
	if err != nil {
		return fmt.Errorf("插入数据失败: %v", err)
	}

	fmt.Printf("失物ID为 %d 的失物信息已成功更新\n", LostItemID)
	return nil
}
