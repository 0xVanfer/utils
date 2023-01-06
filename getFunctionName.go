package utils

import "runtime"

// Get the function name running now.
func RunFuncName() (name string) {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}
