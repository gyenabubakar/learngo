package main

import "fmt"

func main() {
	const (
		ETH = 9 - iota
		BCC
	)

	fmt.Println(ETH)
	fmt.Println(BCC)
}
