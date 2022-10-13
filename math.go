package utils

import "github.com/0xVanfer/types"

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
