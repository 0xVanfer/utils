package utils

import (
	crypto_rand "crypto/rand"
	"math/big"
	math_rand "math/rand"
	"time"
)

// Random from timestamp.
func MathRandBelow(length int) int {
	s := math_rand.NewSource(time.Now().Unix())
	r := math_rand.New(s)
	i := r.Intn(length)
	return i
}

// Real random/ Crypto random.
func CryptoRandBelow(length int) int {
	n, _ := crypto_rand.Int(crypto_rand.Reader, big.NewInt(int64(length)))
	return int(n.Int64())
}

// Choose "times" different numbers between "start" and "end".
func CryptoRandBetween(start int, end int, times int) []int {
	insert := start
	var results []int
	results = append(results, insert)
	for i := 0; i < times; i++ {
		// larger the interval
		bigInt := big.NewInt(int64(end - insert))
		n, _ := crypto_rand.Int(crypto_rand.Reader, bigInt)
		insert += int(n.Int64()) + 1
		if insert >= end {
			return CryptoRandBetween(start, end, times)
		}
		results = append(results, insert)
	}
	results = append(results, end)
	return results
}
