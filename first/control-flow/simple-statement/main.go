package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	convSuccess = "Converted int successfully. %q is %d.\n"
	convFailed  = "Conversion failed. Pass a valid int. :("
	invalidArgs = "Pass in the right args."
)

func main() {
	args := os.Args

	if n := len(args); n < 2 {
		fmt.Println(invalidArgs)
	} else if convNumber, err := strconv.Atoi(args[1]); err != nil {
		fmt.Println(convFailed)
	} else {
		fmt.Printf(convSuccess, args[1], convNumber)
	}
}
