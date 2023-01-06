package utils

import (
	"math/big"
)

// Sqrt for big int.
//
// Examples:
//
//	BigSqrt(100, 2) = 10
//	BigSqrt(1000, 3) = 10
func BigSqrt(num *big.Int, times int) *big.Int {
	if times < 2 {
		return num
	}
	var n, a, b, m, m2 big.Int
	// n = num
	n = *num
	// a = 1
	a.SetInt64(int64(1))
	// b = n = num
	b.Set(&n)

	for {
		m.Add(&a, &b).Div(&m, big.NewInt(2))
		if m.Cmp(&a) == 0 || m.Cmp(&b) == 0 {
			break
		}
		// m2 = m^times
		m2.Mul(&m, &m)
		for i := 0; i < times-2; i++ {
			m2.Mul(&m2, &m)
		}
		if m2.Cmp(&n) > 0 {
			b.Set(&m)
		} else {
			a.Set(&m)
		}
	}
	return &m
}
