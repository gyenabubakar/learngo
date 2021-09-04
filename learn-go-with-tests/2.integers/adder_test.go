package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	actual := Adder(1,2)
	expected := 3

	if actual != expected {
		t.Errorf("Expected %d. Received %d.\n", expected, actual)
	}
}

func ExampleAdder() {
	sum := Adder(2,4)
	fmt.Println(sum)
	// Output: 6
}