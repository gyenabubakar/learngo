package main

func Sum(nums [5]int) int {
	var sum int
	for _, n := range nums {
		sum += n
	}
	return sum
}
