package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Gyen")
	wanted := "Hello Gyen!"

	if got != wanted {
		t.Errorf("got %q; wanted %q\n", got, wanted)
	}
}
