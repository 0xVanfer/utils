package utils

import (
	"fmt"
	"time"
)

// Recover from panic and restart.
func Restart(functionToRecover func(), functionName string) {
	if r := recover(); r != nil {
		fmt.Println(r)
		fmt.Println(TimeNowString(), functionName, "recovered from panic, will restart.")
		functionToRecover()
	}
}

// Recover from panic and restart after waiting for a while.
func RestartAndSleep(functionToRecover func(), functionName string, sleepTime time.Duration) {
	if r := recover(); r != nil {
		fmt.Println(r)
		fmt.Println(TimeNowString(), functionName, "recovered from panic, will restart.")
		time.Sleep(sleepTime)
		functionToRecover()
	}
}
