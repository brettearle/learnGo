package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Brett")
	want := "Hello, Brett"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}