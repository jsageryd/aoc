package main

import "testing"

func TestCountOverlaps(t *testing.T) {
	makeLine := func(x1, y1, x2, y2 int) line {
		return line{
			from: coord{x: x1, y: y1},
			to:   coord{x: x2, y: y2},
		}
	}

	lines := []line{
		makeLine(0, 9, 5, 9),
		makeLine(8, 0, 0, 8),
		makeLine(9, 4, 3, 4),
		makeLine(2, 2, 2, 1),
		makeLine(7, 0, 7, 4),
		makeLine(6, 4, 2, 0),
		makeLine(0, 9, 2, 9),
		makeLine(3, 4, 1, 4),
		makeLine(0, 0, 8, 8),
		makeLine(5, 5, 8, 2),
	}

	if got, want := countOverlaps(lines), 5; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
