package db

import (
	"fmt"
)

func AddApplication(relateItemId int, sender, receiver string, submitDate int) error {
	// 准备 SQL 语句
	sqlStmt := `INSERT INTO tab_applications (lost_item_id, submitter_address, receiver, found_date)
                VALUES (?, ?, ?, ?)`

	// 执行 SQL 语句插入数据
	_, err := Db.Exec(sqlStmt, relateItemId, sender, receiver, submitDate)
	if err != nil {
		return fmt.Errorf("插入数据失败: %v", err)
	}

	return nil
}
