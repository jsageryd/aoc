package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := []string{
		"..@@.@@@@.",
		"@@@.@.@.@@",
		"@@@@@.@.@@",
		"@.@@@@..@.",
		"@@.@@@@.@@",
		".@@@@@@@.@",
		".@.@.@.@@@",
		"@.@@@.@@@@",
		".@@@@@@@@.",
		"@.@.@@@.@.",
	}

	if got, want := part1(input), 13; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
