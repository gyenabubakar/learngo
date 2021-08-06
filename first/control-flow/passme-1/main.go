package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: [username] [password]")
		return
	}

	var (
		username, password string = "gyen", "1234"
		uInput, pInput     string = os.Args[1], os.Args[2]
	)

	if username == uInput && pInput == password {
		fmt.Printf("Access granted, %s.\n", username)
	} else {
		fmt.Printf("Access denied for %s.\n", username)
	}
}
