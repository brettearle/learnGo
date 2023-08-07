package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello with first name", func(t *testing.T) {
		got := Hello("Brett", "")
		want := "Hello, Brett"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello with empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello with spanish", func(t *testing.T) {
		got := Hello("Brett", "Spanish")
		want := "Hola, Brett"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello with french", func(t *testing.T) {
		got := Hello("Brett", "French")
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
