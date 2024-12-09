package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := []byte("2333133121414131402\n")

	if got, want := part1(input), 1928; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
