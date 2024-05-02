package eth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"io/ioutil"
	"strings"
)

// ParseTransaction 解析交易数据
func ParseTransaction(tx *types.Transaction, abiFilePath string) (string, []interface{}, error) {
	// 读取 ABI 文件
	abiBytes, err := ioutil.ReadFile(abiFilePath)
	if err != nil {
		return "", nil, fmt.Errorf("读取 ABI 文件失败: %v", err)
	}

	// 创建合约 ABI 对象
	contractAbi, err := abi.JSON(strings.NewReader(string(abiBytes)))
	if err != nil {
		return "", nil, fmt.Errorf("创建合约 ABI 失败: %v", err)
	}

	// 获取交易数据
	data := tx.Data()
	fmt.Println(data)
	// 解析交易数据
	if len(data) > 0 {
		// 获取合约方法的签名
		methodSig := data[:4]

		// 根据方法签名查找方法
		method, err := contractAbi.MethodById(methodSig)
		if err != nil {
			return "", nil, fmt.Errorf("查找方法失败: %v", err)
		}

		// 解析参数
		args, err := method.Inputs.UnpackValues(data[4:])
		if err != nil {
			return "", nil, fmt.Errorf("解析参数失败: %v", err)
		}

		// 返回方法名和参数
		return method.Name, args, nil
	}

	// 如果交易数据为空，则返回空方法名和空参数列表
	return "", nil, nil
}
