package main

import "testing"

var groups = [][]string{
	{"abc"},
	{"a", "b", "c"},
	{"ab", "ac"},
	{"a", "a", "a", "a"},
	{"b"},
}

func TestPart1(t *testing.T) {
	if got, want := part1(groups), 11; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := part2(groups), 6; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
