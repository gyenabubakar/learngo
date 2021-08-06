package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}

	// numbers
	fmt.Printf("%f \n", 12.5)
	fmt.Printf("%.2f \n", 12.5256)
	fmt.Printf("%e \n", 123400000.0)
	fmt.Printf("%E \n", 123400000.0)

	fmt.Println()

	// values, structs
	fmt.Printf("%v \n", p)
	fmt.Printf("%+v \n", p)
	fmt.Printf("%#v \n", p)

	fmt.Println()

	// number systems, code points
	fmt.Printf("%b \n", 14)
	fmt.Printf("%c \n", 25) // prints character corresponding to codepoint
	fmt.Printf("%x \n", 15) // prints hexadecimal
	fmt.Printf("%x\n", "hello")

	fmt.Println()

	// pointers
	fmt.Printf("%p \n", &p)
	fmt.Printf("%T \n", &p)

	fmt.Println()

	// width/padding
	fmt.Printf("|%5d| \n", 45)
	fmt.Printf("|%5d| \n", 45)
	fmt.Printf("|%-5d| \n", 45)
	fmt.Printf("|%-7.2f| \n", 45.567)
	fmt.Printf("|%-6s|%-6s| \n", "hello", "world")
}
