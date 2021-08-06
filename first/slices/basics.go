package main

import (
	"fmt"
	"strconv"
	"strings"
	// s "github.com/inancgumus/prettyslice"
)

const dailyExpenses string = `10 239 34
12 34
33 42 24
12 35 94 23`

func main() {
	expenses, expensesError := fetch(dailyExpenses)

	if expensesError != nil {
		fmt.Println("Error occured when fetching expenses. Try again later.")
		return
	}

	for i, expense := range expenses {
		fmt.Printf("Day %d: Amount spent is - GHS %d\n", i+1, sumOf(expense))
	}
}

// Accepts a string of numbers seperated by newlines and transforms them into a multi-dimensional slice of int slices containing the prices as elements.
func fetch(expenses string) ([][]int, error) {
	trimmedExpenses := strings.Split(expenses, "\n")
	expensesSlice := make([][]int, 0, len(trimmedExpenses))

	for _, day := range trimmedExpenses {
		strPrices := strings.Fields(day)
		amounts := make([]int, 0, len(strPrices))

		for _, price := range strPrices {
			priceInt, convError := strconv.Atoi(price)

			if convError != nil {
				fmt.Println("The passed expenses string has an invalid int literal.")
				return nil, convError
			}

			amounts = append(amounts, priceInt)
		}

		expensesSlice = append(expensesSlice, amounts)
	}

	return expensesSlice, nil
}

func sumOf(prices []int) int {
	var sum int

	for _, price := range prices {
		sum += price
	}

	return sum
}

// func br() {
// 	fmt.Println()
// 	fmt.Println()
// }
