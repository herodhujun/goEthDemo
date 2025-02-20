package clientinit

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func Initclient() ethclient.Client {
	url := "https://sepolia.infura.io/v3/2f61b797a8ae452da9ed145f0b2ba9ed" //
	//url :="https://cloudflare-eth.com";
	client, err := ethclient.Dial(url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	return *client // we'll use this in the upcoming sections
}
