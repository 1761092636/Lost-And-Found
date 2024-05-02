package eth

import (
	"fmt"
	"math/big"
	"strconv"
)

func GetHistory() (arr []History, err error) {

	//调用方法获取总数
	data, err := Ins.TotalHistory(nil)
	if err != nil {
		return
	}
	datastr := data.String()
	date, err := strconv.Atoi(datastr)
	println("totally found ", date, "record")
	if err != nil {
		fmt.Println(err)
		return
	}

	//遍历
	for key := 0; key < date; key++ {
		ans, err := Ins.Histories(nil, big.NewInt(int64(key)))
		arr = append(arr, History(ans))

		fmt.Println("found error have:", err)

	}

	if err != nil {
		fmt.Println(err)
		return
	}

	return arr, nil

}
