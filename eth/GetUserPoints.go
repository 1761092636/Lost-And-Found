package eth

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func GetUserPoints(address common.Address) (proint *big.Int, err error) {

	proint, err = Ins.UserPoints(nil, address)

	return proint, nil

}
