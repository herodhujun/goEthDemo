package contractCode

import (
	"context"
	"encoding/hex"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetContractCode(client ethclient.Client, contractStr string) string {

	contractAddress := common.HexToAddress(contractStr)
	bytecode, err := client.CodeAt(context.Background(), contractAddress, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(bytecode) // 60806...10029
}
