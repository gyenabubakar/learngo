package main

import (
	"fmt"
	"strconv"
)

func main() {
	name, age := "Gyen", 19
	statement := name + " is " + strconv.Itoa(age)
	fmt.Println(statement)
	fmt.Println()

	isMale := false
	statement = "And he's a boy. [" + strconv.FormatBool(isMale) + "]"
	fmt.Println(statement)
}
