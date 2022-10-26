package utils

import (
	"fmt"
)

func ExampleStringToTrueLocations() {
	fmt.Println(StringToTrueLocations("1432351"))
	// Output:
	// [0 1 2 3 4 8 9 11 12 14 15 16 18 20]
}

func ExampleTrueLocationsToString() {
	fmt.Println(TrueLocationsToString([]int64{1, 2, 33}))
	// Output:
	// 8589934598
}
