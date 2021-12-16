package main

import "testing"

func TestCountOverlaps(t *testing.T) {
	lines := []line{
		{0, 9, 5, 9},
		{8, 0, 0, 8},
		{9, 4, 3, 4},
		{2, 2, 2, 1},
		{7, 0, 7, 4},
		{6, 4, 2, 0},
		{0, 9, 2, 9},
		{3, 4, 1, 4},
		{0, 0, 8, 8},
		{5, 5, 8, 2},
	}

	t.Run("Without diagonals", func(t *testing.T) {
		if got, want := countOverlaps(lines, false), 5; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("With diagonals", func(t *testing.T) {
		if got, want := countOverlaps(lines, true), 12; got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
