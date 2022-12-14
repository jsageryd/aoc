package main

import "testing"

var input = []string{
	"498,4 -> 498,6 -> 496,6",
	"503,4 -> 502,4 -> 502,9 -> 494,9",
}

func TestPart1(t *testing.T) {
	if got, want := part1(input), 24; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCave_Step(t *testing.T) {
	cave := newCave(input)

	for cave.step() {
	}

	wantCaveStr := `
......+...
..........
......o...
.....ooo..
....#ooo##
...o#ooo#.
..###ooo#.
....oooo#.
.o.ooooo#.
#########.`[1:]

	if got, want := cave.String(), wantCaveStr; got != want {
		t.Errorf("got:\n%s\n\nwant:\n%s", got, want)
	}
}
