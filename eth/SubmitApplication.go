package eth

import (
	"Lost_and_Found/db"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"time"
)

func SubmitApplication(opts *bind.TransactOpts, LostItemId *big.Int) error {
	// 调用合约的 AddLostItem 方法
	tx, err := Ins.SubmitFoundApplication(opts, LostItemId)
	if err != nil {
		return err // 返回错误
	}
	owner, err := Ins.GetLostItemOwner(nil, LostItemId)
	if err != nil {
		return err
	}
	// 使用交易详情
	fmt.Println("交易哈希:", tx.Hash().Hex())
	fmt.Println("发送方地址:", opts.From.String())
	fmt.Println("接收方地址:", owner.String())
	fmt.Println("交易时间:", time.Now().Unix())
	var id = LostItemId.Int64()
	var sender = opts.From.String()
	var timestamp = time.Now().Unix()
	// 其他交易详情信息
	err = db.AddApplication(int(id), sender, owner.String(), int(timestamp))
	if err != nil {
		return err
	}
	fmt.Println("提交申请成功")

	return nil
}
