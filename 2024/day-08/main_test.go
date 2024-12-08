package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := []string{
		"............",
		"........0...",
		".....0......",
		".......0....",
		"....0.......",
		"......A.....",
		"............",
		"............",
		"........A...",
		".........A..",
		"............",
		"............",
	}

	if got, want := part1(input), 14; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
