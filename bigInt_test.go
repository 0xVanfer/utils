package utils

import (
	"math/big"
	"testing"
)

func TestBigInt(t *testing.T) {
	BigSqrt(big.NewInt(4096), 6)
}
