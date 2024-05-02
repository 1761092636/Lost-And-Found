package eth

import (
	"Lost_and_Found/db"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"sort"
)

// Rank 函数用于计算用户排名并返回排名信息列表
func Rank() []RankInfo {
	// 获取所有用户信息
	err, users := db.GetAllUser()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// 存储排名信息的列表
	var rankList []RankInfo

	// 遍历所有用户
	for _, user := range users {
		// 获取用户积分信息
		points, err := Ins.UserPoints(nil, common.HexToAddress(user.Address))
		if err != nil {
			println(err)
			continue
		}

		// 将用户信息、积分信息存储到排名信息列表中
		rankList = append(rankList, RankInfo{
			User:   user,
			Points: points,
		})
	}

	// 对排名信息列表根据积分进行排序
	sort.Slice(rankList, func(i, j int) bool {
		return rankList[i].Points.Cmp(rankList[j].Points) > 0
	})

	// 给排名信息列表中的每个用户添加排名信息
	for i, rankInfo := range rankList {
		rankInfo.Rank = i + 1
		rankList[i] = rankInfo
	}
	println(rankList)
	return rankList
}
