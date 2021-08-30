package main

import "fmt"

type author struct {
	name, nationality string
}
type book struct {
	title, date, name string
	author
}

func main() {
	darmani := author{name: "Lawrence Darmani", nationality: "Ghanaian"}

	griefChild := book{
		author: darmani,
		title: "The Grief Child",
		date: "1995",
	}

	fmt.Printf("%#v\n", griefChild)
	fmt.Printf("%v\n", griefChild.name)
}
