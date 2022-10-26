package utils

import (
	"fmt"
)

func ExampleMin() {
	fmt.Println(Min[int]())
	fmt.Println(Min[float64](1, 2))
	fmt.Println(Min(1, 2, 3, 4, 5))
	// Output:
	// 0
	// 1
	// 1
}

func ExampleMax() {
	fmt.Println(Max[int]())
	fmt.Println(Max[float64](1, 2))
	fmt.Println(Max(1, 2, 3, 4, 5))
	// Output:
	// 0
	// 2
	// 5
}
