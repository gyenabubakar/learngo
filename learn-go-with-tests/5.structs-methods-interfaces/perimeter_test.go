package measurement

import "testing"

func TestPerimeter(t *testing.T) {
	a := Perimeter(10.0, 10.0)
	e := 40.0

	if a != e {
		t.Errorf("Expected %.2f; Received %.2f\n", e, a)
	}
}
