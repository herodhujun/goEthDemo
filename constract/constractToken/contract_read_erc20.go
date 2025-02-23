package constractToken

import (
	"log"

	token "goethdemo/token"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Person struct {
	Name string
	Age  int
}

func GetToken(client *ethclient.Client, constractAddress string) *token.Token {

	// Golem (GNT) Address
	tokenAddress := common.HexToAddress(constractAddress)
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	return instance

	// address := common.HexToAddress(accountAddress)
	// bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// name, err := instance.Name(&bind.CallOpts{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// symbol, err := instance.Symbol(&bind.CallOpts{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// decimals, err := instance.Decimals(&bind.CallOpts{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("name: %s\n", name)         // "name: Golem Network"
	// fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
	// fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"

	// fmt.Printf("wei: %s\n", bal) // "wei: 74605500647408739782407023"

	// fbal := new(big.Float)
	// fbal.SetString(bal.String())
	// value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

	// fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
}
