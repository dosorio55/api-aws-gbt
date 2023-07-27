package tools

import (
	"fmt"
	"time"
)

// GetTimeNow returns the current time
func GetTimeNow() string {
	t := time.Now()
	// the "d" is from diggit, and we are formating it
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	// return time.Now().Format("2006-01-02 15:04:05")
}
