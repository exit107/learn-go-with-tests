package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if want != got {
			t.Errorf("wanted: %q, got: %q", want, got)
		}
	}

	t.Run("say 'Hello, Dave' when 'Dave' is passed as a string", func(t *testing.T) {
		want := "Hello, Dave"
		got := Hello("Dave", "English")
		assertCorrectMessage(t, got, want)

	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		want := "Hello, World"
		got := Hello("", "English")
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}
