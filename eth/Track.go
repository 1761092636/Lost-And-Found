package eth

import (
	"fmt"
)

// Track 获取溯源信息
func Track(traceHash [32]byte) (TraceInfo, error) {

	// 生成 traceHash
	fmt.Println(traceHash)
	// 调用合约的 getTraceInfo 函数
	traceInfo, err := Ins.GetTraceInfo(nil, traceHash)
	if err != nil {
		return TraceInfo{}, err
	}

	return TraceInfo(traceInfo), nil
}
