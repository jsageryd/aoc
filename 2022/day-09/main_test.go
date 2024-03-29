package main

import "testing"

func TestMove(t *testing.T) {
	for n, tc := range []struct {
		input     []string
		knotCount int
		want      int
	}{
		{[]string{"R 4", "U 4", "L 3", "D 1", "R 4", "D 1", "L 5", "R 2"}, 2, 13},
		{[]string{"R 4", "U 4", "L 3", "D 1", "R 4", "D 1", "L 5", "R 2"}, 10, 1},
		{[]string{"R 5", "U 8", "L 8", "D 3", "R 17", "D 10", "L 25", "U 20"}, 10, 36},
	} {
		if got, want := move(tc.input, tc.knotCount), tc.want; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
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

func TestChebyshevDistance(t *testing.T) {
	for n, tc := range []struct {
		a, b     coord
		distance int
	}{
		{coord{0, 0}, coord{0, 0}, 0},

		{coord{0, 0}, coord{0, -1}, 1},
		{coord{0, 0}, coord{-1, -1}, 1},
		{coord{0, 0}, coord{-1, 0}, 1},
		{coord{0, 0}, coord{-1, 1}, 1},
		{coord{0, 0}, coord{0, 1}, 1},
		{coord{0, 0}, coord{1, 1}, 1},
		{coord{0, 0}, coord{1, 0}, 1},
		{coord{0, 0}, coord{1, -1}, 1},

		{coord{0, 0}, coord{0, -2}, 2},
		{coord{0, 0}, coord{-1, -2}, 2},
		{coord{0, 0}, coord{-2, -2}, 2},
		{coord{0, 0}, coord{-2, -1}, 2},
		{coord{0, 0}, coord{-2, 0}, 2},
		{coord{0, 0}, coord{-2, 1}, 2},
		{coord{0, 0}, coord{-2, 2}, 2},
		{coord{0, 0}, coord{-1, 2}, 2},
		{coord{0, 0}, coord{0, 2}, 2},
		{coord{0, 0}, coord{1, 2}, 2},
		{coord{0, 0}, coord{2, 2}, 2},
		{coord{0, 0}, coord{2, 1}, 2},
		{coord{0, 0}, coord{2, 0}, 2},
		{coord{0, 0}, coord{2, -1}, 2},
		{coord{0, 0}, coord{2, -2}, 2},
		{coord{0, 0}, coord{1, -2}, 2},
	} {
		if got, want := chebyshevDistance(tc.a, tc.b), tc.distance; got != want {
			t.Errorf("[%d] chebyshevDistance(%v, %v) = %d, want %d", n, tc.a, tc.b, got, want)
		}
	}
}
