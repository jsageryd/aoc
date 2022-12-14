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

func TestPart2(t *testing.T) {
	if got, want := part2(input), 93; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCave_Step(t *testing.T) {
	t.Run("Without floor", func(t *testing.T) {
		cave := newCave(input, false)

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
	})

	t.Run("With floor", func(t *testing.T) {
		cave := newCave(input, true)

		for cave.step() {
		}

		wantCaveStr := `
............o............
...........ooo...........
..........ooooo..........
.........ooooooo.........
........oo#ooo##o........
.......ooo#ooo#ooo.......
......oo###ooo#oooo......
.....oooo.oooo#ooooo.....
....oooooooooo#oooooo....
...ooo#########ooooooo...
..ooooo.......ooooooooo..
#########################`[1:]

		if got, want := cave.String(), wantCaveStr; got != want {
			t.Errorf("got:\n%s\n\nwant:\n%s", got, want)
		}
	})
}
