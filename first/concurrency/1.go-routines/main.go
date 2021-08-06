package main

import (
	"fmt"
	// "runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// runtime.GOMAXPROCS(1)
	start := time.Now()

	wg.Add(2)
	fmt.Println("Start!")

	// lowercase letters
	go func() {

		defer wg.Done()

		for i := 0; i < 3; i++ {
			for c := 'a'; c <= 'z'; c++ {
				fmt.Print(string(c))
			}
			fmt.Println()
		}

		fmt.Println()
		fmt.Println()
	}()

	// uppercase letters
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			for c := 'A'; c <= 'Z'; c++ {
				fmt.Print(string(c))
			}
			fmt.Println()
		}
		fmt.Println()
		fmt.Println()
	}()

	wg.Wait()
	fmt.Println("Done!")

	fmt.Println()
	fmt.Println()

	elapsed := time.Since(start)
	fmt.Println("Time taken to run:", elapsed.Milliseconds(), "ms.")
}
