package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Pass a string as argument!")
		return
	}

	const (
		link     = "http://"
		linkSize = len(link)
	)

	var (
		text = args[0]
		size = len(text)
		buff = make([]byte, 0, size)
	)

	for i := 0; i < size; i++ {
		buff = append(buff, text[i])
		//fmt.Println(string(buff))

		if len(text[i:]) >= linkSize && text[i:linkSize+i] == link {
			fmt.Printf("text[%d:%d + %[1]d] = %[3]v\n", i, linkSize, text[i:linkSize+i])
		}
	}
}
