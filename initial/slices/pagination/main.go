package main

import (
	"fmt"
	"math"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	names := []string{
		"gyen", "saul", "charles", "sena",
		"bernard", "adwoa", "peter", "lizzy",
		"kofi", "philip", "rose", "lydia",
		"biney", "cash",
	}

	s.MaxPerLine = 4

	const pageSize = 4
	pageNumber := 1

	nNames := len(names)

	for initIndex := 0; initIndex < nNames; initIndex += pageSize {
		limit := initIndex + pageSize

		var currentPage []string

		if limit > nNames {
			currentPage = names[initIndex:]
		} else {
			currentPage = names[initIndex:limit]
		}

		pageTitle := fmt.Sprintf(
			"Page %d of %d",
			pageNumber,
			int(math.Ceil(float64(nNames)/float64(pageSize))),
		)
		s.Show(pageTitle, currentPage)

		pageNumber++
	}
}
