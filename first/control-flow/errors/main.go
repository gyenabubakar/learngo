package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	number, _error := strconv.Atoi(os.Args[1])
	if _error != nil {
		fmt.Println("Please enter a real number. 👇")
		fmt.Println(_error)
		return
	}

	fmt.Printf("Converted string, %q, is now %d\n", os.Args[1], number)
}
