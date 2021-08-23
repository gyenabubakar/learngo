package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	baton := make(chan int)

	wg.Add(1)

	go startRace(baton)

	// start race
	baton <- 1

	wg.Wait()
}

func startRace(baton chan int) {
	rand.Seed(time.Now().UnixNano())

	cRunner := <-baton
	var nRunner int

	fmt.Printf("Runner %d running...\n", cRunner)

	if cRunner != 4 {
		nRunner = cRunner + 1
		fmt.Printf("Runner %d getting set!\n", nRunner)
		go startRace(baton)
	}

	// generate a randing time for running
	rn := rand.Intn(5)
	time.Sleep(time.Second * time.Duration(rn))
	fmt.Printf("Runner %d took %d seconds.\n", cRunner, rn)

	// end race if current runner is the last
	if cRunner == 4 {
		fmt.Printf("Runner %d finished; race over!\n", cRunner)
		wg.Done()
		return
	}

	fmt.Printf("Runner %d exchange to Runner %d.\n", cRunner, nRunner)
	fmt.Println()
	baton <- nRunner
}
