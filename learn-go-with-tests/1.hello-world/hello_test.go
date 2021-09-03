package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("return 'Hello NAME!' when called with none-empty string", func(t *testing.T) {
		actual := Hello("Gyen")
		expected := "Hello Gyen!"

		if actual != expected {
			t.Errorf("Expected %q. Received %q.\n", expected, actual)
		}
	})

	t.Run("return 'Hello World!' if called with empty string", func(t *testing.T) {
		actual := Hello("")
		expected := "Hello World!"

		if actual != expected {
			t.Errorf("Expected %q. Received %q.\n", expected, actual)
		}
	})
}
