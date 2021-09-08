package measurement

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rectangle{Width: 10, Height: 10}
	a := Perimeter(r)
	e := 40.0

	if a != e {
		t.Errorf("Expected %.2f; Received %.2f\n", e, a)
	}
}
