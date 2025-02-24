package main

import (
	"context"
	"fmt"
	"goethdemo/address"
	"goethdemo/clientinit"
	"goethdemo/constract/constractExcute"

	// "goethdemo/contractCode"

	"log"
	"math"
	"math/big"
)

func main() {
	//初始化客户端
	client := clientinit.Initclient()
	//生成帐户地址
	address := address.GetAddr()
	//查询最新区块账号余额
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance) // 25893180161173005034
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 25.729324269165216041

	//获取网络上最新区块信息
	//transaction.GetTransaction(&client)
	// transaction.TransETH(client, "c6673dc1aad4c7848dd8836fcf8636b37e13dab4aafdd7eb58b60606e097ddeb", "0x0754942A85dF6D15219D51B9d3BF56C0e92CC560")
	//部署测试合约
	// constractDeploy.Deploy(client, "c6673dc1aad4c7848dd8836fcf8636b37e13dab4aafdd7eb58b60606e097ddeb")
	// webUrl := "wss://sepolia.drpc.org"
	// constractUrl := "0x02213E3d90b755339d846278992C4FcE111D7704"
	// subscribeEvent.Event(webUrl, constractUrl)
	constractExcute.Excute(client, "c6673dc1aad4c7848dd8836fcf8636b37e13dab4aafdd7eb58b60606e097ddeb", "0x02213E3d90b755339d846278992C4FcE111D7704")
	//获取erc20合约实例
	/**
	erc20Token := constractToken.GetToken(&client, "0x44499312f493F62f2DFd3C6435Ca3603EbFCeeBa")
	accountAddressStr := "0x3409e0738662C3F5d8eC4ff780bea60C945f8FB2"
	accountAddress := common.HexToAddress(accountAddressStr)
	bal, err := erc20Token.BalanceOf(&bind.CallOpts{}, accountAddress)
	if err != nil {
		log.Fatal(err)
	}

	name, err := erc20Token.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	symbol, err := erc20Token.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	decimals, err := erc20Token.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"

	fmt.Printf("wei: %s\n", bal) // "wei: 74605500647408739782407023"

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

	fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
	**/

	//关闭客户端
	defer client.Close()
}
