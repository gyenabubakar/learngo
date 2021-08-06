package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	inputs := os.Args[1:]

	if len(inputs) != 2 {
		fmt.Println("[Your name] [pos|neg]")
		return
	}

	moods := [2][4]string{
		{"happy! 😁", "terrific! 😎", "satisfied! 😋", "okay. 🙃"},
		{"sad! 😥", "depressed! 😪", "hungry! 😫", "terrible 🤮"},
	}

	var moodType [4]string

	if inputs[1] == "pos" {
		moodType = moods[0]
	} else {
		moodType = moods[1]
	}

	rand.Seed(time.Now().UnixNano())

	moodIndex := rand.Intn(len(moodType))

	currentMood := moodType[moodIndex]

	fmt.Printf("%s feels %s\n", inputs[0], currentMood)
}
