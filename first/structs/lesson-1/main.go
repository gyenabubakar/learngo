package main

import "fmt"

func main() {
	type person struct {
		firstname string
		lastname  string
		age       int
	}

	gyen := person{
		firstname: "Abubakar",
		lastname: "Sadik",
	}

	var sena person
	sena.age = 21
	sena.firstname = "Godsway"

	fmt.Printf("%#v\n", gyen)
	fmt.Println()
	fmt.Printf("%#v\n", sena)
}
