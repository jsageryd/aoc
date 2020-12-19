package main

import "testing"

func TestPart1(t *testing.T) {
	var input = []string{
		"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		"mem[8] = 11",
		"mem[7] = 101",
		"mem[8] = 0",
	}

	if got, want := part1(input), 165; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
