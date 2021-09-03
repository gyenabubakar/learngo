package main

import "testing"

func TestHello(t *testing.T) {
	assert := func(t testing.TB, e, a string) {
		t.Helper()
		if a != e {
			t.Errorf("Expected %q. Received %q.\n", e, a)
		}
	}

	t.Run("return 'Hello NAME!' when called with none-empty string", func(t *testing.T) {
		actual := Hello("Gyen", "")
		expected := "Hello Gyen!"

		assert(t, expected, actual)
	})

	t.Run("return 'Hello World!' if called with empty string", func(t *testing.T) {
		actual := Hello("", "")
		expected := "Hello World!"

		assert(t, expected, actual)
	})

	t.Run("greet in Spanish", func(t *testing.T) {
		a := Hello("Gyen", "Spanish")
		e := "Hola Gyen!"

		assert(t, e, a)
	})

	t.Run("greet in French", func(t *testing.T) {
		e := "Bonjour Gyen!"
		a := Hello("Gyen", "French")

		assert(t, e, a)
	})
}
