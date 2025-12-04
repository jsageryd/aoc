package main

import (
	"bytes"
	"testing"
)

var input = []string{
	"987654321111111",
	"811111111111119",
	"234234234234278",
	"818181911112111",
}

func TestPart1(t *testing.T) {
	if got, want := part1(input), 357; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := part2(input), 3121910778619; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestMaxComb(t *testing.T) {
	digits := []byte("154356")
	k := 3

	gotComb := maxComb(digits, k)
	wantComb := []byte("556")

	if !bytes.Equal(gotComb, wantComb) {
		t.Errorf("got %q, want %q", gotComb, wantComb)
	}
}
