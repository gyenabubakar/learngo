package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var start, end rune

	if args := os.Args[1:]; len(args) == 2 {
		start = []rune(args[0])[0]
		end = []rune(args[1])[0]
	} else {
		start, end = 'A', 'Z'
	}

	fmt.Printf(
		"%-10s %-10s %-10s %-10s \n",
		"Char", "Dec", "Hex", "Unicode",
	)
	fmt.Println(strings.Repeat("-", 48))

	for codepoint := start; codepoint <= end; codepoint++ {
		fmt.Printf("%-10[1]c| %-10[1]d| %-10[1]x| % -12x\n", codepoint, string(codepoint))
		fmt.Println(strings.Repeat("-", 48))
	}
}
