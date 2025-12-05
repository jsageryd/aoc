package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := []string{
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

	if got, want := part1(input), 3; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
