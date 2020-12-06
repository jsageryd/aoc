package main

import "testing"

func TestPart1(t *testing.T) {
	groups := [][]string{
		{"abc"},
		{"a", "b", "c"},
		{"ab", "ac"},
		{"a", "a", "a", "a"},
		{"b"},
	}

	if got, want := part1(groups), 11; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
