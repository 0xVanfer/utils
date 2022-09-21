package utils

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/math"
)

// Consider a map: map[loc]=bool.
// Convert the coded string into locations.
func StringToTrueLocations(str string) []int64 {
	// string to big int
	x := big.NewInt(0)
	big, _ := x.SetString(str, 10)
	// big int to bytes
	bs := big.Bytes()
	// bytes to binary string
	buf := bytes.NewBuffer([]byte{})
	for _, v := range bs {
		buf.WriteString(fmt.Sprintf("%08b", v))
	}
	binStr := buf.String()
	// binary string to locations
	var newLocs []int64
	for i := 0; i < len(binStr); i++ {
		if binStr[len(binStr)-i-1:len(binStr)-i] == "1" {
			newLocs = append(newLocs, int64(i))
		}
	}
	return newLocs
}

// Consider a map: map[loc]=bool.
// Use a bigInt to locate all true locations and convert into string.
func TrueLocationsToString(trueLocs []int64) string {
	// should not repeat
	newLocs := RemoveRepetitionInArray(trueLocs)
	newBig := big.NewInt(0)
	for _, i := range newLocs {
		newBig.Add(newBig, math.BigPow(2, i))
	}
	return newBig.String()
}
