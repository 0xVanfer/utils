package utils

import (
	"github.com/ethereum/go-ethereum/crypto"
)

// Get the hash of a string.
//
// Example:
//
//	GetStringHash("TraderJoeSupplyEvent(address,address,uint256,address)") = "0x222da511b15d564e69f7e3832eaefbe7fb9875c94fbeecdc358917f0fd0a4f9f"
func GetStringHash(input string) string {
	hashRes := crypto.Keccak256Hash([]byte(input))
	return hashRes.Hex()
}

// Get hash and use it as private key, and generate an eth address.
//
// Return hash(private key) and eth address.
//
// Example:
//
//	GetHashAsPriv("8") = "0xe4b1702d9298fee62dfeccc57d322a463ad55ca201256d01f62b45b2e1c21c10", "0xe0FC04FA2d34a66B779fd5CEe748268032a146c0"
func GetHashAsPriv(input string) (string, string) {
	hashRes := crypto.Keccak256Hash([]byte(input))
	hashStr := hashRes.Hex()
	address := Private2Address(hashStr[2:])
	return hashStr, address
}
