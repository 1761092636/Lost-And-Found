package eth

import "github.com/ethereum/go-ethereum/common"

func GetManger() common.Address {
	address, err := Ins.GetAdminAddress(nil)
	if err != nil {
		return address
	}
	return address
}
