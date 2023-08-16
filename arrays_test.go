package utils

import (
	"fmt"
	"testing"
)

func TestArrayAnd(t *testing.T) {
	fmt.Println(ArrayAnd([]int{1, 2, 2, 3, 4}, []int{2, 3, 5, 8}))
	fmt.Println(ArrayOr([]int{1, 2, 2, 3, 4}, []int{2, 3, 5, 8}))
}
