// Naming a package as main makes it executable
package main

/*
Importing packages:
Imports are scoped to only this file.
This is a shorthand for importing several packages.
*/
import (
	"fmt"
	"runtime"
)

// main function is the entry point for this executable package
func main() {
	var cpus = runtime.NumCPU()
	fmt.Printf("You have %v CPUs!", cpus)
}
