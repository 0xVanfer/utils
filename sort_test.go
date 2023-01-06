package utils

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	// list := []int{1, 2, 5, 6, 3, 12}
	fmt.Println(SortSimple(true, []int{1, 2, 5, 6, 3, 12}))
}

func TestSortMap(t *testing.T) {
	mapp := map[string]int{
		"a": 1,
		"b": 3,
		"c": 4,
		"d": 2,
	}
	x, y := SortSimpleMap(false, mapp)
	fmt.Println(x)
	fmt.Println(y)
}
