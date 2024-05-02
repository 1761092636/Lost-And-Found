package eth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func GetHistoryByUser(user common.Address) (arr []History, err error) {

	records, err := Ins.GetHistoryRecordsByOperator(&bind.CallOpts{}, user)
	if err != nil {
		return nil, fmt.Errorf("failed to get history records: %v", err)
	}

	// 将 Solidity 返回的记录数组转换为 Go 中的结构体数组
	var historyRecords []History
	for _, record := range records {
		historyRecords = append(historyRecords, History{
			HistoryID:         record.HistoryID,
			Operator:          record.Operator,
			Operation:         record.Operation,
			OperationTime:     record.OperationTime,
			RelatedLostItemID: record.RelatedLostItemID,
		})
	}

	return historyRecords, nil
}
