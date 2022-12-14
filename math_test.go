package utils

import (
	"fmt"
)

func ExampleMin() {
	fmt.Println(Min[int]())
	fmt.Println(Min[float64](1, 2))
	fmt.Println(Min(1, 2, 3, 4, 5))
	fmt.Println(Min([]int{2, 3, 4}...))
	// Output:
	// 0
	// 1
	// 1
	// 2
}

func ExampleMax() {
	fmt.Println(Max[int]())
	fmt.Println(Max[float64](1, 2))
	fmt.Println(Max(1, 2, 3, 4, 5))
	fmt.Println(Max([]int{2, 3, 4}...))
	// Output:
	// 0
	// 2
	// 5
	// 4
}

func ExampleMedium() {
	list := []int{1, 2, 5, 4, 3}
	fmt.Println(Medium(list...))

	list = []int{1, 2, 5, 4, 3, 2}
	fmt.Println(Medium(list...))
	// Output:
	// 3
	// 2.5
}

func ExampleMode() {
	list := []int{1, 2, 5, 4, 3, 2}
	fmt.Println(Mode(list...))
	list = []int{1, 2, 5, 4, 3, 2, 3, 4}
	fmt.Println(Mode(list...))
	// Output:
	// [2]
	// [2 3 4]
}
