package main

import (
	"context"
	"fmt"
	"goethdemo/address"
	"goethdemo/clientinit"
	"log"
	"math"
	"math/big"
)

func main() {
	client := clientinit.Initclient()
	address := address.GetAddr()
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance) // 25893180161173005034
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	header, _ := client.HeaderByNumber(context.Background(), nil)
	fmt.Println("Go 代码连接的节点最新区块号:", header.Number)
	fmt.Println(ethValue) // 25.729324269165216041
}
