package main

import (
	"fmt"
)

func main() {
	structs()
}

type house struct {
	name  string
	rooms int
}

func structs() {
	myHouse := house{
		name:  "Gyen's Villa",
		rooms: 5,
	}

	addRoom(myHouse)

	fmt.Println()
	fmt.Println()

	fmt.Printf("structs -> %p --> %+v\n", &myHouse, myHouse)

	addRoomPtr(&myHouse)

	fmt.Println()
	fmt.Println()

	fmt.Printf("structs -> %p --> %+v\n", &myHouse, myHouse)
}

func addRoom(h house) {
	fmt.Println()
	fmt.Println()

	h.rooms++
	fmt.Printf("addRoom -> %p --> %+v\n", &h, h)
}

func addRoomPtr(h *house) {
	fmt.Println()
	fmt.Println()

	h.rooms++
	fmt.Printf("addRoomPtr -> %[1]p --> %+[1]v\n", h)

}
