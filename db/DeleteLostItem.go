package db

import (
	"fmt"
)

// DeleteLostItem 从数据库中删除指定失物ID的失物项
func DeleteLostItem(lostItemID int) error {
	// 准备SQL语句
	query := "DELETE FROM tab_lostitems WHERE lostItemID = ?"

	// 准备并执行数据库操作
	stmt, err := Db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// 执行删除操作
	_, err = stmt.Exec(lostItemID)
	if err != nil {
		return err
	}

	fmt.Printf("失物ID为 %d 的失物已成功删除\n", lostItemID)
	return nil
}
