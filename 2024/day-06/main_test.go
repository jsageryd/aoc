package main

import (
	"testing"
)

var input = []string{
	"....#.....",
	".........#",
	"..........",
	"..#.......",
	".......#..",
	"..........",
	".#..^.....",
	"........#.",
	"#.........",
	"......#...",
}

func TestPart1(t *testing.T) {
	if got, want := part1(input), 41; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := part2(input), 6; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
