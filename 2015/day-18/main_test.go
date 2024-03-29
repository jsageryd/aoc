package main

import (
	"strings"
	"testing"
)

var input = []string{
	".#.#.#",
	"...##.",
	"#....#",
	"..#...",
	"#.#..#",
	"####..",
}

func TestStep(t *testing.T) {
	grid := parse(input, false)

	step(grid, 4, false)

	gotStr := draw(grid)

	wantStr := `
......
......
..##..
..##..
......
......`[1:]

	if gotStr != wantStr {
		t.Errorf("got\n%s\nwant\n%s", gotStr, wantStr)
	}
}

func TestCountOn(t *testing.T) {
	grid := parse(input, false)

	if got, want := countOn(grid), 15; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestParseAndDraw(t *testing.T) {
	gotStr := draw(parse(input, false))
	wantStr := strings.Join(input, "\n")

	if gotStr != wantStr {
		t.Errorf("got\n%s\nwant\n%s", gotStr, wantStr)
	}
}
