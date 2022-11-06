package main

import (
	"fmt"
	"testing"
)

func TestRepeatString(t *testing.T) {
	got := RepeatString("a", 5)
	exp := "aaaaa"

	if got != exp {
		t.Errorf("expected %q, got %q", exp, got)
	}
}

func BenchmarkRepeatString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatString("j", 5)
	}
}

func ExampleRepeatString() {
	got := RepeatString("B", 5)
	fmt.Println(got)
	// Output: BBBBB
}
