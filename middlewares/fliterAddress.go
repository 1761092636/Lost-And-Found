package middlewares

import "github.com/ethereum/go-ethereum/common"

// FilterEmptyAddresses 从地址数组中过滤掉空地址
func FilterEmptyAddresses(addresses []common.Address) []common.Address {
	var filteredAddresses []common.Address
	emptyAddress := common.HexToAddress("0x0000000000000000000000000000000000000000")
	for _, address := range addresses {
		if address.String() != emptyAddress.String() {
			filteredAddresses = append(filteredAddresses, address)
		}
	}
	return filteredAddresses
}
