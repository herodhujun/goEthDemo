package transaction

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetTransaction(client *ethclient.Client) {
	//获取最新区块信息
	block, error := client.BlockByNumber(context.Background(), nil)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("最新区块号:", block.Number) // 5671744
	blockHashHex := block.Hash().Hex()
	fmt.Println("最新区块hash:", blockHashHex) // 5671744
	//获取链id
	// chainID, err := client.NetworkID(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println("当前链id:", block.Number) // 5671744
	/**
	for _, tx := range block.Transactions() {
		fmt.Println("交易hash:", tx.Hash().Hex())            // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		fmt.Println("交易value:", tx.Value().String())       // 10000000000000000
		fmt.Println("交易gas:", tx.Gas())                    // 105000
		fmt.Println("交易gasPrice:", tx.GasPrice().Uint64()) // 102000000000
		fmt.Println("交易nonce:", tx.Nonce())                // 110644
		fmt.Println("交易data:", tx.Data())                  // []
		fmt.Println("交易to:", tx.To().Hex())

		//获取交易发送方
		if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
			fmt.Println("交易from:", sender.Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		}
		//获取交易对应的收据
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		//交易收据状态
		fmt.Println("交易receipt.Status:", receipt.Status) // 1

	}
	**/

	blockHash := common.HexToHash(blockHashHex)
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("交易数量:", count) // 1
	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
	}
}
