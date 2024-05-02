package eth

import (
	"fmt"
	"math/big"
)

func GetLostItemById(key *big.Int) (arr []LostItem, err error) {

	ans, err := Ins.LostItems(nil, key)
	arr = append(arr, LostItem(ans))

	if err != nil {
		fmt.Println(err)
		return
	}

	return arr, nil

}
