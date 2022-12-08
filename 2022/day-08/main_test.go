package main

import "testing"

func TestPart1(t *testing.T) {
	input := []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}

	if got, want := part1(input), 21; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestVisible(t *testing.T) {
	input := []string{
		"30373",
		"25512",
		"65332",
		"33549",
		"35390",
	}

	for n, tc := range []struct {
		x, y    int
		visible bool
	}{
		{0, 0, true},
		{4, 0, true},
		{0, 4, true},
		{4, 4, true},
		{1, 1, true},
		{2, 1, true},
		{3, 1, false},
		{1, 2, true},
		{2, 2, false},
		{3, 2, true},
		{1, 3, false},
		{2, 3, true},
		{3, 3, false},
	} {
		if got, want := visible(input, tc.x, tc.y), tc.visible; got != want {
			t.Errorf("[%d] (%d, %d) got %t, want %t", n, tc.x, tc.y, got, want)
		}
	}
}
