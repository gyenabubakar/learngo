package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("sum collection of any size", func(t *testing.T) {
		n := []int{1, 2}
		actual := Sum(n)
		expected := 3

		if actual != expected {
			t.Errorf("Expected %d. Received %d. Given %#v\n", expected, actual, n)
		}
	})
}
