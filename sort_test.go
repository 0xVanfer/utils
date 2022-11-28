package utils

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	list := []int{1, 2, 5, 6, 3, 2, 1, 1, 1, 1, 1, 1, 11, 12, 12, 12, 12, 2}
	x := SortSimple(true, list)
	fmt.Println(x)
}
