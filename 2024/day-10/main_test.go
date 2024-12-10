package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := []string{
		"89010123",
		"78121874",
		"87430965",
		"96549874",
		"45678903",
		"32019012",
		"01329801",
		"10456732",
	}

	if got, want := part1(input), 36; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
