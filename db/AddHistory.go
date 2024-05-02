package db

import "fmt"

func AddHistory(RelatedLostItemId int, Operator, Operation string, OperationTime int) error {
	// 准备 SQL 语句
	sqlStmt := `INSERT INTO tab_history (RelatedLostItemId, Operator, Operation, OperationTime)
                VALUES (?, ?, ?, ?)`

	// 执行 SQL 语句插入数据
	_, err := Db.Exec(sqlStmt, RelatedLostItemId, Operator, Operation, OperationTime)
	if err != nil {
		return fmt.Errorf("插入数据失败: %v", err)
	}

	return nil
}
