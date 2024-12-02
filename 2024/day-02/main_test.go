package main

import "testing"

var input = []string{
	"7 6 4 2 1",
	"1 2 7 8 9",
	"9 7 6 2 1",
	"1 3 2 4 5",
	"8 6 4 4 1",
	"1 3 6 7 9",
}

func TestPart1(t *testing.T) {
	if got, want := part1(input), 2; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := part2(input), 4; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
