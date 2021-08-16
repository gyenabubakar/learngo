package main

import "fmt"

type minute int64
type second int64

func main() {
	// type Duration int64
	// var timeTakenMS Duration = 3600
	// fmt.Printf("%T", timeTakenMS)

	var (
		oneHour          minute = 60
		oneHourInSeconds second
	)

	oneHourInSeconds = second(oneHour) * 60

	fmt.Printf("There are %d seconds in 1 hour.\n", oneHourInSeconds)
}
