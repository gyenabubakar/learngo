package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	randomsMax, argsErr := strconv.Atoi(os.Args[1:][0])

	if argsErr != nil {
		fmt.Println("You passed an invalid integer as argument. :(")
		return
	}

	var randoms []int
	var uniques []int

	for index := 0; index < randomsMax; index++ {
		randomNumber := rand.Intn(randomsMax)
		randoms = append(randoms, randomNumber)
	}

	sort.Ints(randoms)

	fmt.Println("Random numbers:", randoms)

	for index, val := range randoms {
		if index > 0 {
			if uniques[len(uniques)-1] == val {
				continue
			}

			uniques = append(uniques, val)
			continue
		}

		uniques = append(uniques, val)
	}

	fmt.Println("Unique numbers:", uniques)
}
