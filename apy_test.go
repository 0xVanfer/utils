package utils

import (
	"fmt"
)

func ExampleApr2Apy() {
	fmt.Println(Apr2Apy(0.1))
	fmt.Println(Apr2Apy(1))
	// Output:
	// 0.10515578161622718
	// 1.7145674820220145
}

func ExampleApy2Apr() {
	fmt.Println(Apy2Apr(0.1))
	fmt.Println(Apy2Apr(1))
	// Output:
	// 0.09532262476476205
	// 0.693805752190747
}
