package main

import (
	"fmt"
)

func main() {
	var engToTwi = map[string]string{
		"come":  "bra",
		"eat":   "didi",
		"dance": "sa",
		"look":  "hwɛ",
	}

	fmt.Printf("%#v\n", engToTwi)

	delete(engToTwi, "dance")

	fmt.Printf("%#v\n", engToTwi)

	for key := range engToTwi {
		delete(engToTwi, key)
	}

	fmt.Printf("%#v\n", engToTwi)
}
