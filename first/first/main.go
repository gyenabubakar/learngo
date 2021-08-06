package main

// import "fmt"

func main() {
	const (
		mon = iota + 1
		tue = iota + 1
		wed = iota + 1
		thu = iota + 1
		fri = iota + 1
		sat = iota + 1
		sun = iota + 1
	)

	println(mon)
	println(tue)
	println(wed)
	println(thu)
	println(fri)
	println(sat)
	println(sun)
}
