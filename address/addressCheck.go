package address

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// AddressType 地址类型枚举
type AddressType string

const (
	ContractAddress AddressType = "contract"
	AccountAddress  AddressType = "account"
	InvalidAddress  AddressType = "invalid"
)

// GetAddressType 判断地址类型
func GetAddressType(ctx context.Context, client ethclient.Client, addressStr string) (AddressType, error) {
	// 1. 验证地址格式
	if !common.IsHexAddress(addressStr) {
		return InvalidAddress, fmt.Errorf("invalid Ethereum address")
	}

	// 3. 转换为Address类型
	address := common.HexToAddress(addressStr)

	// 4. 查询代码哈希
	code, err := client.CodeAt(ctx, address, nil) // nil表示最新区块
	if err != nil {
		return InvalidAddress, fmt.Errorf("failed to get code: %v", err)
	}

	// 5. 判断代码长度
	if len(code) > 0 {
		return ContractAddress, nil
	}
	return AccountAddress, nil
}
