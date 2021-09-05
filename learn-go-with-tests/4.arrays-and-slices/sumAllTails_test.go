package main

import "testing"

func TestSumAllTails(t *testing.T) {
	assert := func(tb testing.TB, expected, actual []int) {
		t.Helper()
		if !IntSlicesAreEqual(expected, actual) {
			tb.Errorf("Expected %v. Received %v.\n", expected, actual)
		}
	}

	t.Run("sum tails of some slices", func(t *testing.T) {
		a := SumAllTails([]int{1, 2}, []int{3, 4, 5})
		e := []int{2, 9}

		assert(t, e, a)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		a := SumAllTails([]int{}, []int{3, 4, 5})
		e := []int{0, 9}

		assert(t, e, a)
	})
}
