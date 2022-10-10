package utils

import (
	"math"

	"github.com/0xVanfer/types"
)

// Compound times = 365.
//
// Input 0.01 for 1%.
func Apr2Apy[T types.Number](apr T) (apy float64) {
	return math.Pow((1+types.ToFloat64(apr)/365), 365) - 1
}

// Compound times = 365.
//
// Input 0.01 for 1%.
func Apy2Apr[T types.Number](apy T) (apr float64) {
	return (math.Pow(1+types.ToFloat64(apy), 1.0/365) - 1) * 365
}

// Input 0.01 for 1%.
func Apr2ApyWithCompoundTimes[T types.Number, S types.Number](apr T, compoundTimes S) (apy float64) {
	return math.Pow((1+types.ToFloat64(apr)/types.ToFloat64(compoundTimes)), types.ToFloat64(compoundTimes)) - 1
}

// Input 0.01 for 1%.
func Apy2AprWithCompoundTimes[T types.Number, S types.Number](apy T, compoundTimes S) (apr float64) {
	return (math.Pow(1+types.ToFloat64(apy), 1.0/types.ToFloat64(compoundTimes)) - 1) * types.ToFloat64(compoundTimes)
}
