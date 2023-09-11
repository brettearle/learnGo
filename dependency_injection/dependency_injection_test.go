package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Brett")

	got := buffer.String()
	want := "Hello, Brett"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
