package main

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	exp := 4

	if sum != exp {
		t.Errorf("expected %d got %d", exp, sum)
	}
}
