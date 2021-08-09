package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// get user input in terminal
	integers, err := getUserNumbers()
	// log any errors and terminate program
	if err != nil {
		fmt.Println(err)
		return
	}

	// channel for synchronising goroutines
	channel := make(chan []int)

	// empty 12-element capacity int slice, to hold all ints from all 4 goroutines
	allInts := make([]int, 0, 12)

	// initialise goroutine 4 times
	for i := 0; i < 4; i++ {
		go sortInts(integers[i], channel)
	}

	// wait for messages from 4 goroutines and append them to allInts
	for i := 0; i < 4; i++ {
		n := <- channel
		allInts = append(allInts, n...)
	}

	fmt.Println()

	// sort all ints
	sort.Ints(allInts)
	fmt.Println("All ints sorted:", allInts)
}

//4,12,56,6,11,45,76,10,74,3,14,23

// getUserNumbers
func getUserNumbers() ([][]int, error) {
	fmt.Println("Enter 12 numbers, separated with a comma (,).")
	fmt.Println("NB: Do not leave a trailing comma. Let's go! 👇")
	fmt.Println()
	fmt.Printf("👉 ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	fmt.Println()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	r, _ := regexp.Compile("([0-9]{1,2},){11}([0-9]{1,2}){1}")

	if r.MatchString(input) {
		// convert string to rune slice
		r := []rune(input)

		// l (limit)
		// detach trailing comma and newline characters
		l := len(r) - 2

		// convert to string
		input = string(r[:l])

		// split input using comma
		numsStr := strings.Split(input, ",")

		// make an empty int slice
		integers := make([]int, 0, len(numsStr))

		// convert string slice into int slice
		// fill results in "integers"
		for _, el := range numsStr {
			num, err := strconv.Atoi(el)
			if err != nil {
				fmt.Println(err)
				return nil, err
			}

			integers = append(integers, num)
		}

		// split slice into 4 parts
		slices := [][]int{integers[:3], integers[3:6], integers[6:9], integers[9:12]}

		return slices, nil
	}

	return nil, errors.New("your input was invalid. Ensure you entered numbers separated with commas")
}

// sortInts takes an int slice, sorts it and passes it through channel
func sortInts(n []int, channel chan []int) {
	sort.Ints(n)
	fmt.Println("Sub-array:", n)
	channel <- n
}