package utils

import (
	"fmt"

	"github.com/0xVanfer/types"
)

// Simple sort. Do not need to define Len(), Less(), Swap().
//
// Should not be used to sort too much amount of values.
// If the length is over 12, sort.Sort() suggested.
//
// `ascending`: True: The result will be ascending.
// False: The result will be descending.
//
// Example:
//
//	SortSimple(true, []int{1, 2, 5, 6, 3, 12}...) = []int{1, 2, 3, 5, 6, 12}
func SortSimple[T types.Ordered](ascending bool, input ...T) []T {
	length := len(input)
	if length == 0 {
		return []T{}
	}
	if length > 100 {
		fmt.Println("Input too long, sort.Sort() suggested.")
	}
	if length > 1000 {
		return []T{}
	}
	// sort.insertionSort
	for i := 1; i < length; i++ {
		if ascending {
			// Ascending.
			for j := i; j > 0 && input[j] < input[j-1]; j-- {
				input[j-1], input[j] = input[j], input[j-1]
			}
		} else {
			// Descending.
			for j := i; j > 0 && input[j] > input[j-1]; j-- {
				input[j-1], input[j] = input[j], input[j-1]
			}
		}
	}
	return input
}

// Simple sort by map value. Do not need to define Len(), Less(), Swap().
//
// Should not be used to sort too much amount of values.
// If the length is over 12, sort.Sort() suggested.
//
// Returns the ordered keys and ordered values.
//
// Example:
//
//	mapp := map[string]int{
//		"a": 1,
//		"b": 3,
//		"c": 4,
//		"d": 2,
//	}
//	x, y := SortSimpleMap(true, mapp)
//	x = []string{"a", "d", "b", "c"}
//	y = []int{1, 2, 3, 4}
func SortSimpleMap[S comparable, T types.Ordered](ascending bool, input map[S]T) ([]S, []T) {
	length := len(input)
	if length == 0 {
		return nil, nil
	}
	if length > 100 {
		fmt.Println("Input too long, sort.Sort() suggested.")
	}
	if length > 1000 {
		return nil, nil
	}

	var xList []S
	var yList []T
	for x, y := range input {
		xList = append(xList, x)
		yList = append(yList, y)
	}

	// sort.insertionSort
	for i := 1; i < length; i++ {
		if ascending {
			// Ascending.
			for j := i; j > 0 && yList[j] < yList[j-1]; j-- {
				yList[j-1], yList[j] = yList[j], yList[j-1]
				xList[j-1], xList[j] = xList[j], xList[j-1]
			}
		} else {
			// Descending.
			for j := i; j > 0 && yList[j] > yList[j-1]; j-- {
				yList[j-1], yList[j] = yList[j], yList[j-1]
				xList[j-1], xList[j] = xList[j], xList[j-1]
			}
		}
	}
	return xList, yList
}
