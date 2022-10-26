package utils

import "github.com/0xVanfer/types"

// Get the minimum one of the input numbers.
//
// Example:
//
//	Min(1, 2, 3) = 1
//	Min[int]()   = 0
//
// NOTE:
//
//	Can not use array as input param.
func Min[T types.OrderedNumber](numbers ...T) T {
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	var min T = numbers[0]
	for _, number := range numbers {
		if number < min {
			min = number
		}
	}
	return min
}

// Get the maximum one of the input numbers.
//
// Example:
//
//	Max(1, 2, 3) = 3
//	Max[int]()   = 0
//
// NOTE:
//
//	Can not use array as input param.
func Max[T types.OrderedNumber](numbers ...T) T {
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers) == 1 {
		return numbers[0]
	}
	var max T = numbers[0]
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}
