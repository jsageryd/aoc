package main

import "testing"

func TestPart1(t *testing.T) {
	input := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}

	if got, want := part1(input), 142; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
