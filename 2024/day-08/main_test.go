package main

import (
	"testing"
)

var input = []string{
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

func TestPart1(t *testing.T) {
	if got, want := part1(input), 14; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := part2(input), 34; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
