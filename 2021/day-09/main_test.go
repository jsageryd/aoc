package main

import "testing"

func TestPart1(t *testing.T) {
	input := []string{
		"2199943210",
		"3987894921",
		"9856789892",
		"8767896789",
		"9899965678",
	}

	if got, want := part1(input), 15; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
