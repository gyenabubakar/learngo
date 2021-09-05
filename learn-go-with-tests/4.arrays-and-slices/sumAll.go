package main

import "sort"

func SumAll(slices ...[]int) []int {
	results := make([]int, 0, len(slices))

	for _, slice := range slices {
		results = append(results, Sum(slice))
	}

	return results
}

// check if two slices are equal
func IntSlicesAreEqual(slice1, slice2 []int) bool {
	// if lengths of slices aren't equal, they're not the same
	if len(slice1) != len(slice2) {
		return false
	}

	// else, sort both slices (increasing order)
	sort.Ints(slice1)
	sort.Ints(slice2)

	// if an element in slice1 isn't equal to an element
	// (in the same index) in slice2, slices aren't equal
	for i, v := range slice1 {
		if slice2[i] != v {
			return false
		}
	}

	// if execution reaches this point, both slices are equal
	return true
}
