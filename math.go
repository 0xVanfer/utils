package utils

import "github.com/0xVanfer/types"

// Get the minimum one of the input numbers.
//
// Example:
//
//	Min(1, 2, 3) = 1
//	Min[int]()   = 0
//	Min([]int{2, 3, 4}...) = 2
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
//	Max([]int{2, 3, 4}...) = 4
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

// Get the medium number.
func Medium[T types.OrderedNumber](numbers ...T) float64 {
	numbers = SortSimple(true, numbers...)
	if len(numbers) == 0 {
		return 0
	}
	if len(numbers)%2 == 1 {
		return types.ToFloat64(numbers[len(numbers)/2])
	} else {
		return (types.ToFloat64(numbers[len(numbers)/2-1]) + types.ToFloat64(numbers[len(numbers)/2])) / 2
	}
}

// Get the mode(math).
func Mode[T types.OrderedNumber](numbers ...T) []T {
	if len(numbers) <= 1 {
		return numbers
	}
	mapp := make(map[T]int64)
	for _, num := range numbers {
		if mapp[num] == 0 {
			mapp[num] = 1
		} else {
			mapp[num] += 1
		}
	}
	numberList, timesList := SortSimpleMap(false, mapp)
	if len(timesList) == 0 {
		return []T{}
	}
	maxTime := timesList[0]
	var res []T
	for i := range numberList {
		if timesList[i] == maxTime {
			res = append(res, numberList[i])
		}
	}
	return SortSimple(true, res...)
}
