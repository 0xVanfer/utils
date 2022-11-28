package utils

import (
	"fmt"

	"github.com/0xVanfer/types"
)

// Simple sort. Don't need to define Len(), Less(), Swap().
//
// Should not be used to sort too much amount of values.
// If the length is over 12, sort.Sort() suggested.
func SortSimple[T types.Ordered](ascending bool, input []T) []T {
	length := len(input)
	if length == 0 {
		return input
	}
	if length > 20 {
		fmt.Println("Input too long, sort.Sort() suggested.")
	}
	if length > 100 {
		return input
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
