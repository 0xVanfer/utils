package utils

import (
	"fmt"
	"runtime"
	"time"
)

// Recover from panic and restart.
//
// NOTE:
//
//	Param `functionName` is no longer in use. It will be read directly from runtime.
func Restart(functionToRecover func(), functionName string) {
	if r := recover(); r != nil {
		// Get the running function.
		pc := make([]uintptr, 2)
		runtime.Callers(2, pc)
		f := runtime.FuncForPC(pc[1])
		fmt.Println(r)
		fmt.Println(TimeNowString(), f.Name(), "recovered from panic, will restart.")
		// At least sleep for 1 sec, in case of endless loop.
		time.Sleep(time.Second)
		functionToRecover()
	}
}

// Recover from panic and restart after waiting for a while.
//
// NOTE:
//
//	Param `functionName` is no longer in use. It will be read directly from runtime.
func RestartAndSleep(functionToRecover func(), functionName string, sleepTime time.Duration) {
	if r := recover(); r != nil {
		// Get the running function.
		pc := make([]uintptr, 2)
		runtime.Callers(2, pc)
		f := runtime.FuncForPC(pc[1])
		fmt.Println(r)
		fmt.Println(TimeNowString(), f.Name(), "recovered from panic, will restart.")
		if sleepTime.Seconds() < 1 {
			sleepTime = time.Second
		}
		time.Sleep(sleepTime)
		functionToRecover()
	}
}
