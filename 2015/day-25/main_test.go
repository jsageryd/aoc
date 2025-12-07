package main

import "testing"

func TestPart1(t *testing.T) {
	input := "row 2, column 1"

	if got, want := part1(input), 31916031; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCoordToIndex(t *testing.T) {
	for n, tc := range []struct {
		x, y  int
		index int
	}{
		{x: 1, y: 1, index: 1},
		{x: 1, y: 2, index: 2},
		{x: 2, y: 1, index: 3},
		{x: 1, y: 3, index: 4},
		{x: 2, y: 2, index: 5},
		{x: 3, y: 1, index: 6},
		{x: 1, y: 4, index: 7},
		{x: 2, y: 3, index: 8},
		{x: 3, y: 2, index: 9},
		{x: 4, y: 1, index: 10},
	} {
		if got, want := coordToIndex(tc.x, tc.y), tc.index; got != want {
			t.Errorf("[%d] coordToIndex(%d, %d) = %d, want %d", n, tc.x, tc.y, got, want)
		}
	}
}
