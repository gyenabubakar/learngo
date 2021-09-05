package main

func SumAllTails(slices ...[]int) []int {
	sumsOfTails := make([]int, 0, len(slices))

	for _, slice := range slices {
		if len(slice) == 0 {
			sumsOfTails = append(sumsOfTails, 0)
			continue
		}

		sumsOfTails = append(sumsOfTails, Sum(slice[1:]))
	}

	return sumsOfTails
}
