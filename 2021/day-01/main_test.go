package main

import "testing"

func TestCountIncreases(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	if got, want := countIncreases(input), 7; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
