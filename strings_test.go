package utils

import (
	"fmt"
)

func ExampleStrUnderline2Seperate() {
	fmt.Println(StrUnderline2Seperate(""))
	fmt.Println(StrUnderline2Seperate("a"))
	fmt.Println(StrUnderline2Seperate("a_b"))
	fmt.Println(StrUnderline2Seperate("a_b_ccc_ddd"))
	// Output:
	//
	// a
	// a b
	// a b ccc ddd
}

func ExampleStrUnderline2SeperateTitle() {
	fmt.Println(StrUnderline2SeperateTitle(""))
	fmt.Println(StrUnderline2SeperateTitle("a"))
	fmt.Println(StrUnderline2SeperateTitle("a_b"))
	fmt.Println(StrUnderline2SeperateTitle("a_b_ccc_ddd"))
	// Output
	//
	// A
	// A B
	// A B Ccc Ddd
}

func ExampleStrUnderline2Camel() {
	fmt.Println(StrUnderline2Camel(""))
	fmt.Println(StrUnderline2Camel("a"))
	fmt.Println(StrUnderline2Camel("a_b"))
	fmt.Println(StrUnderline2Camel("a_b_ccc_ddd"))
	// Output
	//
	// a
	// aB
	// aBCccDdd
}

func ExampleStrCamel2Underline() {
	fmt.Println(StrCamel2Underline("aaBBC"))
	fmt.Println(StrCamel2Underline("A"))
	fmt.Println(StrCamel2Underline("bASccccSS"))
	// Output:
	// aa_b_b_c
	// a
	// b_a_scccc_s_s
}
