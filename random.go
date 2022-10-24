package utils

import (
	crypto_rand "crypto/rand"
	"math/big"
	math_rand "math/rand"
	"sort"
	"time"

	"github.com/0xVanfer/types"
)

// Random from timestamp.
func MathRandBelow[T types.Integer](length T) int {
	s := math_rand.NewSource(time.Now().Unix())
	r := math_rand.New(s)
	i := r.Intn(int(length))
	return i
}

// Real random/ Crypto random.
func CryptoRandBelow[T types.Integer](length T) int {
	n, _ := crypto_rand.Int(crypto_rand.Reader, big.NewInt(int64(length)))
	return int(n.Int64())
}

// Choose "times" different numbers between "start" and "end".
//
// Example:
//
//	CryptoRandBetween(1, 100, 5) = []int{1, 68, 69, 78, 79, 98, 100}
func CryptoRandBetween(start int, end int, times int) []int {
	insert := start
	var results randRes
	results = append(results, insert)
	for i := 0; i < times; i++ {
		new := CryptoRandBelow(end-start) + start
		for ContainInArray(new, results) {
			new = CryptoRandBelow(end-start) + start
		}
		results = append(results, new)
	}
	results = append(results, end)
	sort.Sort(results)
	return results
}

type randRes []int

func (p randRes) Len() int { return len(p) }

func (p randRes) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p randRes) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
