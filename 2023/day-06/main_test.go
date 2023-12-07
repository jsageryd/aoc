package main

import (
	"testing"
)

var input = []string{
	"Time:      7  15   30",
	"Distance:  9  40  200",
}

func TestPart1(t *testing.T) {
	if got, want := part1(input), 288; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := part2(input), 71503; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
