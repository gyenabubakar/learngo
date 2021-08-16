package main

import "fmt"

func main() {
	var (
		width  uint8 = 255
		height       = 255
	)

	fmt.Println(width)
	width++
	fmt.Println(width)
	if int(width) < height {
		fmt.Println("Width is smaller!")
	}
}
