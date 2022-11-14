package utils

import (
	"fmt"
	"testing"
)

func TestCryptoRandFrom(t *testing.T) {
	// range_ := []string{"111", "222", "333", "444", "555"}
	range_ := []int{1, 2, 3, 4, 5, 6, 7}

	for i := 0; i < 10; i++ {
		fmt.Println(CryptoRandFrom(range_))
	}
}
