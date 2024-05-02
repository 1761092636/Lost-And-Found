package eth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

func DeleteLostItems(opts *bind.TransactOpts, LostItemId *big.Int) error {

	tx, err := Ins.DeleteLostItem(opts, LostItemId)
	if err != nil {
		return err // 返回错误
	}

	fmt.Println("交易哈希:", tx.Hash().Hex())

	return nil
}
