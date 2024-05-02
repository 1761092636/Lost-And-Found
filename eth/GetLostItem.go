package eth

import (
	"fmt"
	"math/big"
)

func GetLostItem() ([]LostItem, error) {
	var arr []LostItem
	var err error

	// 假设我们有一个起始ID，这里简化为0，根据实际情况可能有所不同
	startID := big.NewInt(0)

	// 调用方法获取总数
	data, err := Ins.TotalLostItems(nil)
	if err != nil {
		return nil, err
	}

	totalItems := data.Uint64()
	fmt.Println("found", totalItems, "valid lost items")

	// 遍历所有可能的ID，直到找到totalItems个有效失物
	currentID := startID
	for len(arr) < int(totalItems) {
		ans, err := Ins.LostItems(nil, currentID)
		if err != nil {
			// 处理错误，例如打印错误日志
			fmt.Println("Error fetching lost item:", err)
			currentID = currentID.Add(currentID, big.NewInt(1)) // 递增ID并继续下一个循环
			continue
		}

		// 检查User字段是否为空地址，以确认失物是否有效
		if ans.User.String() == "0x0000000000000000000000000000000000000000" {
			// 如果User字段为空，说明该失物已被删除，递增ID并继续下一个循环
			currentID = currentID.Add(currentID, big.NewInt(1))
			continue
		}

		// 将找到的有效失物信息添加到结果数组中
		arr = append(arr, LostItem(ans))

		// 递增ID以查询下一个失物信息
		currentID = currentID.Add(currentID, big.NewInt(1))
	}

	return arr, nil
}
