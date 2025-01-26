package internal

import (
	"time"
)

func GetCurrentTimeString() string {
	currentTime := time.Now()
	return currentTime.Format("20060102150405")
}
