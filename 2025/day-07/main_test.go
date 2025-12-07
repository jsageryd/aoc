package main

import (
	"strings"
	"testing"
)

var input = []string{
	".......S.......",
	"...............",
	".......^.......",
	"...............",
	"......^.^......",
	"...............",
	".....^.^.^.....",
	"...............",
	"....^.^...^....",
	"...............",
	"...^.^...^.^...",
	"...............",
	"..^...^.....^..",
	"...............",
	".^.^.^.^.^...^.",
	"...............",
}

func TestPart1(t *testing.T) {
	if got, want := part1(input), 21; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := part2(input), 40; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestGridString(t *testing.T) {
	grid := parse(input)
	beams := []coord{{1, 1}, {1, 2}, {1, 3}}

	gotString := gridString(grid, beams)

	wantString := strings.TrimSpace(`
.......S.......
.+.............
.+.....^.......
.+.............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
`)

	if gotString != wantString {
		t.Errorf("got:\n%s\n\nwant:\n%s", gotString, wantString)
	}
}
