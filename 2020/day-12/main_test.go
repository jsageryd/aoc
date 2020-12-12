package main

import "testing"

func TestPart1(t *testing.T) {
	input := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}

	if got, want := part1(input), 25; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
