package main

import "fmt"

const engGreeting = "Hello "

func main() {
	fmt.Println(Hello("Gyen"))
}

func Hello(name string) string {
	if name ==  "" {
		return engGreeting + "World!"
	}
	return engGreeting + name + "!"
}
