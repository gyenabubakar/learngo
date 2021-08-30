package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	wanted := "Hello world!"

	if got != wanted {
		t.Errorf("got %q; wanted %q\n", got, wanted)
	}
}
