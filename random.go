package utils

import (
	crypto_rand "crypto/rand"
	"math/big"
	math_rand "math/rand"
	"time"

	"github.com/0xVanfer/types"
)

// Random from timestamp.
//
// Range:
//
//	0 ~ length-1
func MathRandBelow[T types.Integer](length T) int {
	if length == 0 {
		return 0
	}
	s := math_rand.NewSource(time.Now().Unix())
	r := math_rand.New(s)
	i := r.Intn(int(length))
	return i
}

// Real random/ Crypto random.
//
// Range:
//
//	0 ~ length-1
func CryptoRandBelow[T types.Integer](length T) int {
	if length == 0 {
		return 0
	}
	n, _ := crypto_rand.Int(crypto_rand.Reader, big.NewInt(int64(length)))
	return int(n.Int64())
}

// Choose "times" different numbers between "start" and "end".
//
// Example:
//
//	CryptoRandBetween(1, 100, 5) = []int{1, 68, 69, 78, 79, 98, 100}
func CryptoRandBetween(start int, end int, times int) []int {
	if end <= start {
		return []int{}
	}
	insert := start
	var results []int
	results = append(results, insert)
	for i := 0; i < times; i++ {
		new := CryptoRandBelow(end-start) + start
		for ContainInArray(new, results...) {
			new = CryptoRandBelow(end-start) + start
		}
		results = append(results, new)
	}
	results = append(results, end)
	results = SortSimple(true, results...)
	return results
}

// Choose an element randomly from the given array.
func CryptoRandFrom[T any](range_ ...T) T {
	return range_[CryptoRandBelow(len(range_))]
}
