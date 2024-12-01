package main

import "testing"

var input = []string{
	"3   4",
	"4   3",
	"2   5",
	"1   3",
	"3   9",
	"3   3",
}

func TestPart1(t *testing.T) {
	if got, want := part1(input), 11; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := part2(input), 31; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
