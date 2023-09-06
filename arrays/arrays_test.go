package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection 5 numbers sum", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	t.Run("collection dynamic length 3 passed numbers sum", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := SumSlice(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	})

	t.Run("collection dynamic length 6 passed numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		got := SumSlice(numbers)
		want := 21

		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	})
}
