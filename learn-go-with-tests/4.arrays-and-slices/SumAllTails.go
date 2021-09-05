package main

func SumAllTails(slices ...[]int) []int {
	sumsOfTails := make([]int, 0, len(slices))

	for _, slice := range slices {
		sumsOfTails = append(sumsOfTails, Sum(slice[1:]))
	}

	return sumsOfTails
}
