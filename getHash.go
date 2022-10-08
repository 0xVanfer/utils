package utils

import (
	"github.com/ethereum/go-ethereum/crypto"
)

// For example: "TraderJoeSupplyEvent(address,address,uint256,address)".
//
// Space must not be included.
func GetStringHash(input string) string {
	hash := crypto.Keccak256Hash([]byte(input))
	return hash.Hex()
}
