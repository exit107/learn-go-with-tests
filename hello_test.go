package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, Dave"
	got := Hello("Dave")

	if want != got {
		t.Errorf("wanted: %q, got: %q", want, got)
	}
}
