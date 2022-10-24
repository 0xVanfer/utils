package utils

import (
	"crypto/ecdsa"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

// Create a pair of private key and address.
func CreateKey() (privateKey string, address string) {
	private, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(private)
	privateKey = hexutil.Encode(privateKeyBytes)[2:]
	publicKey := private.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return
}

// From private key to address.
func Private2Address(key string) string {
	private, err := crypto.HexToECDSA(key)
	if err != nil {
		return ""
	}
	publicKey := private.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return ""
	}
	return crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
}

// Return checksummed ethereum address.
func ChecksumEthereumAddress(addr string) string {
	hex := strings.ToLower(addr)[2:]
	d := sha3.NewLegacyKeccak256()
	d.Write([]byte(hex))
	hash := d.Sum(nil)
	checksumed := "0x"
	for i, b := range hex {
		c := string(b)
		if b < '0' || b > '9' {
			if hash[i/2]&byte(128-i%2*120) != 0 {
				c = string(b - 32)
			}
		}
		checksumed += c
	}
	return checksumed
}
