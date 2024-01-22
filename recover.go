package utils

import (
	"fmt"
	"runtime"
	"time"
)

// Recover from panic and restart.
func Restart(funcToRecover func()) {
	if r := recover(); r != nil {
		// Get the running function.
		pc := make([]uintptr, 2)
		runtime.Callers(2, pc)
		f := runtime.FuncForPC(pc[1])
		fmt.Println(r)
		fmt.Println(time.Now().UTC().Format("2006-01-02 15:04:05"), f.Name(), "recovered from panic, will restart.")
		// At least sleep for 1 sec, in case of endless loop.
		time.Sleep(time.Second)
		funcToRecover()
	}
}

// Recover from panic and restart after waiting for a while.
func RestartAndSleep(funcToRecover func(), sleepTime time.Duration) {
	if r := recover(); r != nil {
		// Get the running function.
		pc := make([]uintptr, 2)
		runtime.Callers(2, pc)
		f := runtime.FuncForPC(pc[1])
		fmt.Println(r)
		fmt.Println(time.Now().UTC().Format("2006-01-02 15:04:05"), f.Name(), "recovered from panic, will restart.")
		if sleepTime.Seconds() < 1 {
			sleepTime = time.Second
		}
		time.Sleep(sleepTime)
		funcToRecover()
	}
}
