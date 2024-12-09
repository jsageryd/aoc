package main

import (
	"testing"
)

var input = []byte("2333133121414131402\n")

func TestPart1(t *testing.T) {
	if got, want := part1(input), 1928; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := part2(input), 2858; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
