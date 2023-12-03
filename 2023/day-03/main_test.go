package main

import (
	"strings"
	"testing"
)

var input = split(`
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`)

func TestPart1(t *testing.T) {
	if got, want := part1(input), 4361; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := part2(input), 467835; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func split(input string) [][]byte {
	var s [][]byte
	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		s = append(s, []byte(line))
	}
	return s
}
