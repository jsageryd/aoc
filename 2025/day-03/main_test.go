package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := []string{
		"987654321111111",
		"811111111111119",
		"234234234234278",
		"818181911112111",
	}

	if got, want := part1(input), 357; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
