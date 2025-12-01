package main

import "testing"

var input = []string{
	"L68",
	"L30",
	"R48",
	"L5",
	"R60",
	"L55",
	"L1",
	"L99",
	"R14",
	"L82",
}

func TestPart1(t *testing.T) {
	if got, want := part1(input), 3; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := part2(input), 6; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
