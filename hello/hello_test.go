package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("hello with name argument return Hello, name",
		func(t *testing.T) {
			got := Hello("Brett", "")
			want := "Hello, Brett"
			assertCorrectMessage(t, got, want)
		})

	t.Run("called empty returns Hello World",
		func(t *testing.T) {
			got := Hello("", "")
			want := "Hello, world"
			assertCorrectMessage(t, got, want)
		})

	t.Run("called with spanish as prefix returns Hola....",
		func(t *testing.T) {
			got := Hello("Brett", "spanish")
			want := "Hola, Brett"
			assertCorrectMessage(t, got, want)
		})

	t.Run("called with french as prefix returns Bonjour....",
		func(t *testing.T) {
			got := Hello("Brett", "french")
			want := "Bonjour, Brett"
			assertCorrectMessage(t, got, want)
		})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
