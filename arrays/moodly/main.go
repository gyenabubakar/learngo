package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		println("Pass your name!")
		return
	}

	moods := [...]string{
		"happy! 😁",
		"sad! 😥",
		"terrific! 😎",
		"drepressed! 😪",
		"satisfied! 😋",
		"hungry! 😫",
		"okay. 🙃",
	}

	rand.Seed(time.Now().UnixNano())

	currentMood := moods[rand.Intn(len(moods))]

	fmt.Printf("%s feels %s", args[0], currentMood)
}
