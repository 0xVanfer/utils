package utils

import (
	"fmt"
	"testing"
)

func TestRestart(t *testing.T) {
	defer Restart(func() { fmt.Println("test in") }, "")
	panic("err")
}
