package main

import "testing"

func TestSumAll(t *testing.T) {
	a := SumAll([]int{1, 2, 3}, []int{4, 5})
	e := []int{6, 9}

	if !IntSlicesAreEqual(e, a) {
		t.Errorf("Expected %v. Received %v.\n", e, a)
	}
}
