package db

import (
	"fmt"
)

// GetApplicationsByReceiver 根据接收者地址获取所有应用程序记录
func GetApplicationsByReceiver(receiver string) ([]Application, error) {
	// 准备 SQL 查询语句
	sqlStmt := `SELECT id, lost_item_id, submitter_address, receiver, found_date, confirmed, reward_paid, result 
				FROM tab_applications 
				WHERE receiver = ?`

	// 执行 SQL 查询
	rows, err := Db.Query(sqlStmt, receiver)
	if err != nil {
		return nil, fmt.Errorf("执行查询失败: %v", err)
	}
	defer rows.Close()

	// 定义用于存储结果的切片
	var applications []Application

	// 遍历查询结果并映射到结构体中
	for rows.Next() {
		var app Application
		if err := rows.Scan(&app.ApplicationId, &app.RelateItemID, &app.Sender, &app.Receiver, &app.SubmitDate, &app.IsConfirmed, &app.RewardPaid, &app.Result); err != nil {
			return nil, fmt.Errorf("解析结果失败: %v", err)
		}
		applications = append(applications, app)
	}

	// 检查遍历过程中是否出现错误
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("遍历结果时出错: %v", err)
	}

	return applications, nil
}
