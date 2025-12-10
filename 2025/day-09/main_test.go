package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := []string{
		"7,1",
		"11,1",
		"11,7",
		"9,7",
		"9,5",
		"2,5",
		"2,3",
		"7,3",
	}

	if got, want := part1(input), 50; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
