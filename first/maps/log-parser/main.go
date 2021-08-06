package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	webVisitsLog := map[string]int{}

	for input.Scan() {
		fields := strings.Fields(input.Text())
		if len(fields) != 2 {
			fmt.Println("Pass in `Domain Number` format!")
			return
		}

		if num, err := strconv.Atoi(fields[1]); err == nil {
			webVisitsLog[fields[0]] += num
			continue
		} else {
			fmt.Printf("Couldn't %q to int!", fields[1])
			return
		}
	}

	if input.Err() != nil {
		fmt.Println("An error occurred whilst reading input.\n", input.Err())
		return
	}

	fmt.Printf("%-24s  %24s\n", "DOMAIN", "TIMES VISITED")
	fmt.Println(strings.Repeat("-", 50))

	for domain, times := range webVisitsLog {
		fmt.Printf("%-24s  %24d\n", domain, times)
	}

	fmt.Println(strings.Repeat("-", 50))
}