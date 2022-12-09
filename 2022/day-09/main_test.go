package main

import "testing"

func TestPart1(t *testing.T) {
	input := []string{
		"R 4",
		"U 4",
		"L 3",
		"D 1",
		"R 4",
		"D 1",
		"L 5",
		"R 2",
	}

	if got, want := part1(input), 13; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCoord_Move(t *testing.T) {
	for n, tc := range []struct {
		c         coord
		direction string
		want      coord
	}{
		{coord{0, 0}, "L", coord{-1, 0}},
		{coord{0, 0}, "R", coord{1, 0}},
		{coord{0, 0}, "U", coord{0, -1}},
		{coord{0, 0}, "D", coord{0, 1}},
	} {
		tc.c.move(tc.direction)

		if got, want := tc.c, tc.want; got != want {
			t.Errorf("[%d] got %v, want %v", n, got, want)
		}
	}
}

func TestCoord_Follow(t *testing.T) {
	for n, tc := range []struct {
		c     coord
		other coord
		want  coord
	}{
		// Same spot
		{coord{0, 0}, coord{0, 0}, coord{0, 0}},

		// Left/right
		{coord{0, 0}, coord{1, 0}, coord{0, 0}},   // CO  -> CO
		{coord{0, 0}, coord{2, 0}, coord{1, 0}},   // C O -> CO
		{coord{0, 0}, coord{-1, 0}, coord{0, 0}},  // OC  -> OC
		{coord{0, 0}, coord{-2, 0}, coord{-1, 0}}, // O C -> OC

		// Up/down
		{coord{0, 0}, coord{0, 1}, coord{0, 0}},
		{coord{0, 0}, coord{0, 2}, coord{0, 1}},
		{coord{0, 0}, coord{0, -1}, coord{0, 0}},
		{coord{0, 0}, coord{0, -2}, coord{0, -1}},

		// Diagonal
		{coord{0, 0}, coord{1, 1}, coord{0, 0}},
		{coord{0, 0}, coord{2, 2}, coord{1, 1}},
		{coord{0, 0}, coord{-1, -1}, coord{0, 0}},
		{coord{0, 0}, coord{-2, -2}, coord{-1, -1}},

		// Diagonal plus one
		{coord{0, 0}, coord{2, 1}, coord{1, 1}},
		{coord{0, 0}, coord{2, -1}, coord{1, -1}},
		{coord{0, 0}, coord{-2, -1}, coord{-1, -1}},
		{coord{0, 0}, coord{-2, 1}, coord{-1, 1}},
		{coord{0, 0}, coord{1, 2}, coord{1, 1}},
		{coord{0, 0}, coord{-1, 2}, coord{-1, 1}},
		{coord{0, 0}, coord{-1, -2}, coord{-1, -1}},
		{coord{0, 0}, coord{1, -2}, coord{1, -1}},
	} {
		tc.c.follow(tc.other)

		if got, want := tc.c, tc.want; got != want {
			t.Errorf("[%d] got %v, want %v", n, got, want)
		}
	}
}

func TestCoord_Near(t *testing.T) {
	for n, tc := range []struct {
		c     coord
		other coord
		near  bool
	}{
		// Same spot
		{coord{0, 0}, coord{0, 0}, true},

		// Left/right
		{coord{0, 0}, coord{1, 0}, true},
		{coord{0, 0}, coord{2, 0}, false},
		{coord{0, 0}, coord{-1, 0}, true},
		{coord{0, 0}, coord{-2, 0}, false},

		// Up/down
		{coord{0, 0}, coord{0, 1}, true},
		{coord{0, 0}, coord{0, 2}, false},
		{coord{0, 0}, coord{0, -1}, true},
		{coord{0, 0}, coord{0, -2}, false},

		// Diagonal
		{coord{0, 0}, coord{1, 1}, true},
		{coord{0, 0}, coord{2, 2}, false},
		{coord{0, 0}, coord{-1, -1}, true},
		{coord{0, 0}, coord{-2, -2}, false},

		// Diagonal plus one
		{coord{0, 0}, coord{2, 1}, false},
		{coord{0, 0}, coord{2, -1}, false},
		{coord{0, 0}, coord{-2, -1}, false},
		{coord{0, 0}, coord{-2, 1}, false},
		{coord{0, 0}, coord{1, 2}, false},
		{coord{0, 0}, coord{-1, 2}, false},
		{coord{0, 0}, coord{-1, -2}, false},
		{coord{0, 0}, coord{1, -2}, false},
	} {
		if got, want := tc.c.near(tc.other), tc.near; got != want {
			t.Errorf("[%d] %v - %v got %t, want %t", n, tc.c, tc.other, got, want)
		}
	}
}
