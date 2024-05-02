package db

import (
	"strings"
)

func ReverseLookupCategory(category string) (keys []string) {
	for key, cat := range Keywords {
		if cat == category {
			keys = append(keys, key)
		}
	}
	return keys
}
func SelectLostItem(key string) (err error, lostitems []LostItem) {
	var prokeys []string
	moreca, ok := Keywords[strings.ToLower(key)]
	if !ok {
		println("未找到与关键字匹配的类别")

	} else {
		prokeys = ReverseLookupCategory(moreca)
		if len(prokeys) == 0 {
			println("未找到与关键字匹配的类别")
		}
		for _, k := range prokeys {
			println(k)
		}
	}
	var sqlStr = `SELECT * FROM tab_lostitems WHERE itemCategory LIKE ? `

	// 执行查询并遍历prokeys
	for _, k := range prokeys {
		// 执行查询
		rows, err := Db.Query(sqlStr, "%"+k+"%")
		if err != nil {
			return err, nil
		}
		defer rows.Close()

		// 遍历查询结果
		for rows.Next() {
			var lostitem LostItem
			err := rows.Scan(&lostitem.LostItemID, &lostitem.User, &lostitem.FoundOne, &lostitem.LostItemName, &lostitem.Description, &lostitem.LostDate, &lostitem.Location, &lostitem.ItemCategory, &lostitem.IsFound, &lostitem.FoundDate, &lostitem.Reward, &lostitem.LostImg, &lostitem.FoundApplicationsCount)

			if err != nil {
				return err, nil
			}
			lostitems = append(lostitems, lostitem)
		}
	}

	return nil, lostitems

}
