package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// fullName is a package scope variable, that's shared between two Goroutines; call it a shared resource.
	fullName string
	wg       sync.WaitGroup
)

func main() {
	wg.Add(2)

	// concurrently running two function invocations
	go appendName("Godson", "Sena")
	go appendName("Gyen", "Abubakar")

	// wait for both concurrent tasks to complete
	wg.Wait()

	// print the updated value of fullName
	fmt.Println("Full name:", fullName)
}

// function to be ran concurrently
func appendName(first, last string) {
	// schedule wg.Done to be invoked when the function exits
	defer wg.Done()

	// name will be a copy of fullName
	name := fullName

	// if the string is empty
	if name == "" {
		// prepend the combination of first & last to it
		name = first + " " + last
	} else {
		// else, append it to the value it already has
		name = name + " - " + first + " " + last
	}

	// The above if..else blocks perform a READ operation on the variable fullName

	// call for a context switch
	runtime.Gosched()

	// This statement writes to fullName
	fullName = name
}

/*
NOTES:

A Race Condition is encountered with both function invocations:
	➡ appendName("Godson", "Sena")
	➡ appendName("Gyen", "Abubakar")

Each Goroutine reads from the variable fullName, and then writes to it. Hence, there's a communication.
But, is it synchronised? 🤔

Each Goroutine makes a copy of the shared resource, fullName.

This "copy" is now updated depending on its current value - empty or not.
Regardless, there's an assignment; the copy no longer has the previous state of fullName.

Before there's a write to the fullName variable, the current Goroutine
calls the runtime scheduler to switch its context for another Goroutine's.

This new Goroutine starts afresh, reaches line 49 and calls for a context switch as well.
Execution is handed to the previous Goroutine (in our case, since we only have two).

Now, the first Goroutine which then continues to execute line 52, writing to the variable fullName,
and then it exits, since it's done executing.

The scheduler hands executions to the second Goroutine, which also runs its last code, writing to the fullName
variable. Hence, completely overwriting its value.

In the end, we have a race condition because both Go routines perform
unsynchronised read and write operations on a shared resource.
*/
