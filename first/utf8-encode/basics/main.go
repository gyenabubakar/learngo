package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var start, stop int

	if args := os.Args[1:]; len(args) == 2 {
		start, _ = strconv.Atoi(args[0])
		stop, _ = strconv.Atoi(args[1])
	} else {
		start, stop = 'A', 'Z'
	}

	fmt.Printf("%-10s %-10s %-10s %-10s\n", "Literal", "Dec", "Hex", "UTF-8")
	fmt.Println(strings.Repeat("-", 40))

	for i := start; i <= stop; i++ {
		fmt.Printf("%-10c %-10[1]d %-10[1]x % -10x\n", i, string(rune(i)))
	}
}
