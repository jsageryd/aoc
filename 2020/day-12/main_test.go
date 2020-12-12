package main

import "testing"

var input = []string{
	"F10",
	"N3",
	"F7",
	"R90",
	"F11",
}

func TestPart1(t *testing.T) {
	if got, want := part1(input), 25; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := part2(input), 286; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
