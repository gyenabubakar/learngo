package main

import (
	"fmt"
	"os"
)

const (
	usage               = "Usage: [Username] [Password]"
	accessDenied        = "Access denied for %q.\n"
	accessGranted       = "Access granted for %q.\n"
	accessDeniedGeneric = "Access denied for %q."

	u1 = "gyen"
	p1 = "1234"
	u2 = "sena"
	p2 = "abcd"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	if args[1] != u1 && args[1] != u2 {
		fmt.Printf(accessDenied, args[1])
	} else if args[1] == u1 && args[2] == p1 {
		fmt.Printf(accessGranted, u1)
	} else if args[1] == u2 && args[2] == p2 {
		fmt.Printf(accessGranted, u2)
	} else {
		fmt.Printf(accessDeniedGeneric, args[1])
	}
}
