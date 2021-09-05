package main

import "testing"

func TestSumAllTails(t *testing.T) {
	a := SumAllTails([]int{1, 2}, []int{3, 4, 5})
	e := []int{2, 9}

	if !IntSlicesAreEqual(e, a) {
		t.Errorf("Expected %v. Received %v.\n", e, a)
	}
}
