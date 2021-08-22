package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var (
	flag int64
	wg   sync.WaitGroup
)

func main() {
	wg.Add(3)

	startChan := make(chan bool)

	go shutdownTimer(startChan)

	<- startChan
	go do("doing homework", "📚")
	go do("cooking", "🍝")

	wg.Wait()

	fmt.Println("All tasks stopped!")
}

func do(t string, e string) {
	defer wg.Done()

	for {
		if atomic.LoadInt64(&flag) == 1 {
			fmt.Printf("%s Stopped %s.\n\n", e, t)
			break
		}

		fmt.Printf("%s\t%s.\n", e, t)
		time.Sleep(time.Millisecond * 200)
	}
}

func shutdownTimer(sc chan bool) {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(6) + 1
	fmt.Printf("Stopping in %d second(s).\n\n", r)

	sc <- true

	time.Sleep(time.Second * time.Duration(r))

	fmt.Println()
	fmt.Println("Stop all tasks!")
	fmt.Println()

	atomic.StoreInt64(&flag, 1)
	wg.Done()
}
