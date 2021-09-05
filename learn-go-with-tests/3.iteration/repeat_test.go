package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	actual := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if actual != expected {
		t.Errorf("Expected %q. Received %q.", expected, actual)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeatA := Repeat("a", 5)
	fmt.Println(repeatA)
	// Output: aaaaa
}
