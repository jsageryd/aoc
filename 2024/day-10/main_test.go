package main

import (
	"testing"
)

var input = []string{
	"89010123",
	"78121874",
	"87430965",
	"96549874",
	"45678903",
	"32019012",
	"01329801",
	"10456732",
}

func TestPart1(t *testing.T) {
	if got, want := part1(input), 36; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := part2(input), 81; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
