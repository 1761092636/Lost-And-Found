package eth

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
)

func ConfirmApplication(opts *bind.TransactOpts, LostItemId *big.Int, ApplicationId *big.Int, Istrue bool) (string, error) {

	tx, err := Ins.ConfirmFoundApplication(opts, LostItemId, ApplicationId, Istrue)
	if err != nil {
		return "", err // 返回错误
	}

	fmt.Println("交易哈希:", tx.Hash().Hex())
	fmt.Println("审核申请成功")
	ans, _ := Ins.FoundOneChanges(nil, LostItemId)
	TraceHash := hex.EncodeToString(ans.TraceHash[:])
	fmt.Println("TraceHash:", TraceHash)
	return TraceHash, nil
}
