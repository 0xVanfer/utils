package utils

import "github.com/0xVanfer/types"

func Celcius2Fahrenheit[T types.OrderedNumber](degree T) float64 {
	return 1.8*types.ToFloat64(degree) + 32
}

func Fahrenheit2Celcius[T types.OrderedNumber](degree T) float64 {
	return (types.ToFloat64(degree) - 32) / 1.8
}
