package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Specify a directory.")
		return
	}

	fileInfos, dirError := ioutil.ReadDir(args[0])

	if dirError != nil {
		fmt.Println("Oops! No such file or directory!")
		return
	}

	// total capacity for bytes
	var total int

	for _, file := range fileInfos {
		if file.Size() == 0 {
			total += len(file.Name()) + 1
		}
	}

	fileNames := make([]byte, 0, total)

	for _, file := range fileInfos {
		if file.Size() == 0 {
			name := file.Name()

			fileNames = append(fileNames, name...)
			fileNames = append(fileNames, '\n')
		}
	}

	writeError := ioutil.WriteFile("out.txt", fileNames, 0644)

	if writeError != nil {
		fmt.Println(writeError)
		return
	}
}

// func br() {
// 	fmt.Println()
// 	fmt.Println()
// }
