package main

import "testing"

func TestCountIncreases(t *testing.T) {
	for n, tc := range []struct {
		windowSize int
		increases  int
	}{
		{windowSize: 1, increases: 7},
		{windowSize: 3, increases: 5},
	} {
		input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

		if got, want := countIncreases(input, tc.windowSize), tc.increases; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}
