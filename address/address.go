package address

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

func GetAddr() common.Address {
	address := common.HexToAddress("0x164000BCC3716D5487E046e999387f476AE5B29E")

	fmt.Println(address.Hex()) // 0x71C7656EC7ab88b098defB751B7401B5f6d8976F
	// fmt.Println(address.Hash().Hex()) // 0x00000000000000000000000071c7656ec7ab88b098defb751b7401b5f6d8976f
	fmt.Println(address.Bytes()) // [113 199 101 110 199 171 136 176 152 222 251 117 27 116 1 181 246 216 151 111]

	return address
}
