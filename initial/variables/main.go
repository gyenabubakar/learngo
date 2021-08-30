package main

import (
	"fmt"
	"os"
	// "path"
)

func main() {
	// pathToFile := "path/to/file.txt"
	// var dir, file string
	// dir, file = path.Split(pathToFile)

	// fmt.Println("dir :", dir)
	// fmt.Println("file:", file)
	args := os.Args
	fmt.Println(args)
}
