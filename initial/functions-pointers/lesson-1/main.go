package main

import "fmt"

func main() {
	local := 10

	local = limit(local, 5)

	fmt.Println(local)
}

// sets the maximum number n can be through the limit poaram
func limit(n, limit int) (rv int) {
	rv = n
	if rv >= limit {
		rv = limit
	}
	return
}
