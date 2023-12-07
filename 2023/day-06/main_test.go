package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}

	if got, want := part1(input), 288; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
