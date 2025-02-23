package address

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func GetAddr() common.Address {
	address := common.HexToAddress("0x3409e0738662C3F5d8eC4ff780bea60C945f8FB2")

	fmt.Println(address.Hex()) // 0x71C7656EC7ab88b098defB751B7401B5f6d8976F
	// fmt.Println(address.Hash().Hex()) // 0x00000000000000000000000071c7656ec7ab88b098defb751b7401b5f6d8976f
	fmt.Println(address.Bytes()) // [113 199 101 110 199 171 136 176 152 222 251 117 27 116 1 181 246 216 151 111]

	return address
}
