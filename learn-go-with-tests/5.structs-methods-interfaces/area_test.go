package measurement

import "testing"

func TestArea(t *testing.T) {
	a := Area(10, 10)
	e := 100.0

	if a != e {
		t.Errorf("Expected %.2f ; Received %.2f\n", e, a)
	}
}
