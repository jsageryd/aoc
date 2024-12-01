package main

import "testing"

func TestPart1(t *testing.T) {
	input := []string{
		"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}

	if got, want := part1(input), 11; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
