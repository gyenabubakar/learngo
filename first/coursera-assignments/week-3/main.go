package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	integers, err := getUserNumbers()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(integers)
}

//1,2,3,4,5,6,7,8,9,10,11,12

// getUserNumbers
func getUserNumbers() ([][]int, error) {
	fmt.Println("Enter 12 numbers, separated with a comma (,).")
	fmt.Println("NB: Do not leave a trailing comma. Let's go! 👇")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

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
		// fill results in "intergers"
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

// sortInts takes an int slice, sorts it, and returns another int slice
func sortInts(n []int) (r []int) {

}
