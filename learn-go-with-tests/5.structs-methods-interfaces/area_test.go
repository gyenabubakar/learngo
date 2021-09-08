package measurement

import (
	"testing"
)

type TestCase struct {
	name     string
	shape    Shape
	expected float64
}

func TestArea(t *testing.T) {
	assert := func(tb testing.TB, shape Shape, expected float64) {
		tb.Helper()

		actual := shape.Area()

		if actual != expected {
			tb.Errorf("%#v --> Expected %g ; Received %g\n", shape, expected, actual)
		}
	}

	testCases := []TestCase{
		{name: "Rectangle", shape: Rectangle{Width: 10, Height: 7}, expected: 70},
		{name: "Rectangle", shape: Rectangle{Width: 30, Height: 5}, expected: 150},
		{name: "Circle", shape: Circle{Radius: 5}, expected: 78.53981633974483},
		{name: "Circle", shape: Circle{20}, expected: 1256.6370614359173},
	}

	for _, ts := range testCases {
		t.Run(ts.name, func(t *testing.T) {
			assert(t, ts.shape, ts.expected)
		})
	}
}
