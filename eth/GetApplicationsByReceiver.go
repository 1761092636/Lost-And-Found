package eth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
)

// GetApplicationsByReceiver 获取给定接收者地址收到的所有申请
func GetApplicationsByReceiver(receiver common.Address) ([]Application, error) {

	// 获取总失物数
	totalItems, err := Ins.TotalLostItems(nil)
	if err != nil {
		return nil, err
	}

	// 遍历所有失物
	var applications []Application
	var count int64
	var arr []LostItem

	// 假设我们有一个起始ID，这里简化为0，根据实际情况可能有所不同
	startID := big.NewInt(0)

	fmt.Println("found", totalItems, "valid lost items")

	// 遍历所有可能的ID，直到找到totalItems个有效失物
	currentID := startID
	for big.NewInt(int64(len(arr))).Cmp(totalItems) == -1 {
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
		for i := int64(0); i < totalItems.Int64(); i++ {
			// 获取失物ID
			itemID := ans.LostItemID
			// 获取失物信息
			item, err := Ins.LostItems(nil, itemID)
			if err != nil {
				log.Println("Failed to get lost item:", err)
				continue
			}
			var temp = item.FoundApplicationsCount
			count += temp.Int64()

			// 如果失物的接收者是指定的接收者，则获取该失物的所有申请
			if item.User == receiver {
				for j := int64(0); j < item.FoundApplicationsCount.Int64()+count; j++ {
					// 获取申请信息
					app, err := Ins.ApplicationInfo(nil, itemID, big.NewInt(j))
					if err != nil {
						log.Println("Failed to get application:", err)
						continue
					}
					exists, err := Ins.ApplicationExists(nil, itemID, big.NewInt(j))
					if err != nil {
						return nil, err
					}
					if exists == true {
						// 添加申请到列表中

						applications = append(applications, Application{
							ApplicationId: app.ApplicationId,
							RelateItemID:  app.RelateItemID,
							Sender:        app.Mysender,
							Receiver:      app.Receiver,
							SubmitDate:    app.SubmitDate,
							Result:        app.Result,
							IsConfirmed:   app.IsConfirmed,
							RewardPaid:    app.RewardPaid,
						})
					}

				}
				break
			}
		}
		// 递增ID以查询下一个失物信息
		currentID = currentID.Add(currentID, big.NewInt(1))

	}

	return applications, nil
}
