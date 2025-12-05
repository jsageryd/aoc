package main

import (
	"testing"
)

var input = []string{
	"3-5",
	"10-14",
	"16-20",
	"12-18",
	"",
	"1",
	"5",
	"8",
	"11",
	"17",
	"32",
}

func TestPart1(t *testing.T) {
	if got, want := part1(input), 3; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := part2(input), 14; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
