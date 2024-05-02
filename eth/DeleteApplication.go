package eth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

func DeleteApplication(opts *bind.TransactOpts, LostItemId *big.Int, applicationId *big.Int) error {

	tx, err := Ins.DeleteApplication(opts, LostItemId, applicationId)
	if err != nil {
		return err // 返回错误
	}

	fmt.Println("交易哈希:", tx.Hash().Hex())

	fmt.Println("提交申请成功")
	return nil
}
