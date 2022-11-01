package utils

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"strconv"
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

type findAddr struct {
	Private string
	Address string
}

// Find an address contain the target string.
func FindAddressContain(times int, targets ...string) []findAddr {
	for _, target := range targets {
		_, err := strconv.ParseInt(target, 16, 64)
		// not hex
		if err != nil {
			return nil
		}
	}
	var found []findAddr
	var tried int = 0
	for len(found) < times {
		priv, addr := CreateKey()
		if ContainAnyOfStrs(addr, targets) {
			found = append(found, findAddr{Private: priv, Address: addr})
		}
		tried += 1
		if tried > 1000000 {
			fmt.Println("already tried a million times, the target maybe too long")
			return found
		}
	}
	return found
}

// Find an address start with the target string.
func FindAddressStart(times int, targets ...string) []findAddr {
	for _, target := range targets {
		_, err := strconv.ParseInt(target, 16, 64)
		// not hex
		if err != nil {
			return nil
		}
	}
	var found []findAddr
	var tried int = 0
	for len(found) < times {
		priv, addr := CreateKey()
		for _, target := range targets {
			if strings.EqualFold(target, addr[2:len(target)+2]) {
				found = append(found, findAddr{Private: priv, Address: addr})
				continue
			}
		}
		tried += 1
		if tried > 1000000 {
			fmt.Println("already tried a million times, the target maybe too long")
			return found
		}
	}
	return found
}

// Find an address end with the target string.
func FindAddressEnd(times int, targets ...string) []findAddr {
	for _, target := range targets {
		_, err := strconv.ParseInt(target, 16, 64)
		// not hex
		if err != nil {
			return nil
		}
	}
	var found []findAddr
	var tried int = 0
	for len(found) < times {
		priv, addr := CreateKey()
		for _, target := range targets {
			if strings.EqualFold(target, addr[42-len(target):]) {
				found = append(found, findAddr{Private: priv, Address: addr})
				continue
			}
		}
		tried += 1
		if tried > 1000000 {
			fmt.Println("already tried a million times, the target maybe too long")
			return found
		}
	}
	return found
}
