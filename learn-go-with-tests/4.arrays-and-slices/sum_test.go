package main

import "testing"

func TestSum(t *testing.T) {
	n := [...]int{1, 2, 3, 4, 5}
	actual := Sum(n)
	expected := 16

	if actual != expected {
		t.Errorf("Expected %d. Received %d. Given %#v\n", expected, actual, n)
	}
}
