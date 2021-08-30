package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	inputs := os.Args[1:]

	if len(inputs) < 2 {
		fmt.Println("❌ Guess 2 numbers.")
		return
	}
   
	num1, err1 := strconv.Atoi(inputs[0])
	num2, err2 := strconv.Atoi(inputs[1])

	if err1 != nil || err2 != nil {
		if err1 != nil {
			fmt.Printf("❌ %s is not a valid integer.\n", inputs[0])
		}
		if err2 != nil {
			fmt.Printf("❌ %s is not a valid integer.\n", inputs[1])
		}
		return
	}

	rand.Seed(time.Now().UnixNano())
	correctNumber := rand.Intn(11)

	println("💎💎💎 Correct Number:", correctNumber)
	println()
	println("👉👉👉 Your numbers:", num1, num2)
	println()

	if num1 == correctNumber && num2 == correctNumber {
		fmt.Println("✔🔥 SUPER WIN!!! Correct number is", correctNumber)
		fmt.Println("Both of your guessed numbers matched!")
		return
	} else if num1 == correctNumber || num2 == correctNumber {
		fmt.Println("✔🔥 WIN!!! Correct number is", correctNumber)
		fmt.Println("One of your guessed numbers matched!")
		return
	} else {
		fmt.Println("😢 None of your numbers matched.")
		fmt.Println("You lose!")
	}
}
