package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		currentTime time.Time = time.Now()
		tempStr     string    = "Good "
		greeting    string
	)

	switch hour := currentTime.Hour(); {
	case hour >= 0 && hour < 12:
		greeting = tempStr + "morning!"
	case hour > 12 && hour < 17:
		greeting = tempStr + "afternoon!"
	case hour >= 17:
		greeting = tempStr + "evening!"
	}

	fmt.Println(greeting)
}
