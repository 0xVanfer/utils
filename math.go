package utils

import "github.com/0xVanfer/types"

func Min[T types.OrderedNumber](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func Max[T types.OrderedNumber](x, y T) T {
	if x > y {
		return x
	}
	return y
}
