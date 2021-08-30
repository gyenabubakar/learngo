package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	screen "github.com/inancgumus/screen"
)

type digit [5]string

// digits
var (
	zero = digit{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}
	one = digit{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}
	two = digit{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}
	three = digit{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}
	four = digit{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}
	five = digit{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}
	six = digit{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}
	seven = digit{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}
	eight = digit{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}
	nine = digit{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}

	colonOn = digit{
		"   ",
		" █ ",
		"   ",
		" █ ",
		"   ",
	}

	colonOff = digit{
		"   ",
		"   ",
		"   ",
		"   ",
		"   ",
	}
)

func main() {
	// var clock [][]string

	digits := [...]digit{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}

	hrs, mins, secs := getTime()

	showSeparators := true

	clock := make([]digit, 0, 8)

	screen.Clear()
	for {
		// screen.Clear()

		for _, digitStr := range hrs {
			digit, _ := strconv.Atoi(digitStr)
			clock = append(clock, digits[digit])
		}

		if showSeparators {
			clock = append(clock, colonOn)
		} else {
			clock = append(clock, colonOff)
		}

		for _, digitStr := range mins {
			digit, _ := strconv.Atoi(digitStr)
			clock = append(clock, digits[digit])
		}

		if showSeparators {
			clock = append(clock, colonOn)
		} else {
			clock = append(clock, colonOff)
		}

		for _, digitStr := range secs {
			digit, _ := strconv.Atoi(digitStr)
			clock = append(clock, digits[digit])
		}

		for line := 0; line < 5; line++ {
			for _, digit := range clock {
				fmt.Print(digit[line], " ")
			}
			fmt.Println()
		}
		fmt.Println()

		time.Sleep(time.Second)
		showSeparators = !showSeparators
		screen.Clear()
	}
}

func getCurrent(timeType string) []string {
	var slice []string

	switch timeType {
	case "hour":
		slice = strings.Split(strconv.Itoa(time.Now().Hour()), "")
	case "minute":
		slice = strings.Split(strconv.Itoa(time.Now().Minute()), "")
	case "second":
		slice = strings.Split(strconv.Itoa(time.Now().Second()), "")
	}

	digits := make([]string, 0, 2)

	if len(slice) == 1 {
		digits = append(digits, "0", slice[0])
	} else {
		digits = append(digits, slice...)
	}

	return digits
}

type (
	hour   []string
	minute []string
	second []string
)

func getTime() (hour, minute, second) {
	return getCurrent("hour"), getCurrent("minute"), getCurrent("second")
}
