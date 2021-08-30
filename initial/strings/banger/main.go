package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	// get input
	i := os.Args[1]
	// get length os string input
	l := utf8.RuneCountInString(i)
	// banger
	bt := strings.ToUpper(i) + strings.Repeat("!", l)
	fmt.Println(bt)
	fmt.Println(len("eweHi"))
}
