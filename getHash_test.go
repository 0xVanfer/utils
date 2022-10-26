package utils

import (
	"fmt"
)

func ExampleGetHashAsPriv() {
	priv, addr := GetHashAsPriv("100")
	fmt.Println("priv:", priv)
	fmt.Println("addr:", addr)
	// Output:
	// priv: 0x8c18210df0d9514f2d2e5d8ca7c100978219ee80d3968ad850ab5ead208287b3
	// addr: 0xef3CDbFf03CE88F5041C82131b0cF15061b6015E
}
