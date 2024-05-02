package db

import "fmt"

// DeleteApplication 从数据库中删除指定申请
func DeleteApplication(applicationID int) error {
	// 构建 SQL DELETE 语句
	query := "DELETE FROM tab_applications WHERE id = ?"

	// 执行 SQL 语句插入数据
	_, err := Db.Exec(query, applicationID)
	if err != nil {
		return fmt.Errorf("删除数据失败: %v", err)
	}

	return nil
}
