package utils

import (
	"math"

	"github.com/0xVanfer/types"
)

// Compound times = 365.
//
// Example:
//
//	Apr2Apy(0.1) = 0.10515578161622718 // 10%  ==> 10.52%
//	Apr2Apy(1)   = 1.7145674820220145  // 100% ==> 171.46%
func Apr2Apy(apr any) (apy float64) {
	if !types.IsNumber(apr) {
		return 0
	}
	return math.Pow((1+types.ToFloat64(apr)/365), 365) - 1
}

// Example:
//
//	Apr2ApyWithCompoundTimes(0.1, 100) = 0.10511569772075302 // 10%  ==> 10.51%
//	Apr2ApyWithCompoundTimes(1, 100)   = 1.704813829421529   // 100% ==> 170.48%
func Apr2ApyWithCompoundTimes[T types.Number](apr any, compoundTimes T) float64 {
	return math.Pow((1+types.ToFloat64(apr)/types.ToFloat64(compoundTimes)), types.ToFloat64(compoundTimes)) - 1
}

// Compound times = 365.
//
// Example:
//
//	Apy2Apr(0.1) = 0.09532262476476205 // 10%  ==> 9.53%
//	Apy2Apr(1)   = 0.693805752190747   // 100% ==> 69.38%
func Apy2Apr(apy any) (apr float64) {
	if !types.IsNumber(apr) {
		return 0
	}
	return (math.Pow(1+types.ToFloat64(apy), 1.0/365) - 1) * 365
}

// Compound times = 365.
//
// Example:
//
//	Apy2AprWithCompoundTimes(0.1, 100) = 0.09535561438964724 // 10%  ==> 9.54%
//	Apy2AprWithCompoundTimes(1, 100)   = 0.6955550056718884  // 100% ==> 69.56%
func Apy2AprWithCompoundTimes[T types.Number](apy any, compoundTimes T) float64 {
	return (math.Pow(1+types.ToFloat64(apy), 1.0/types.ToFloat64(compoundTimes)) - 1) * types.ToFloat64(compoundTimes)
}
