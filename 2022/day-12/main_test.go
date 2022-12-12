package main

import (
	"fmt"
	"testing"
)

var input = []string{
	"Sabqponm",
	"abcryxxl",
	"accszExk",
	"acctuvwj",
	"abdefghi",
}

func TestPart1(t *testing.T) {
	if got, want := part1(input), 31; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := part2(input), 29; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestAStar(t *testing.T) {
	/*
		s is start
		g is goal
		# is obstacle
		right edge has high cost

		s.....
		###...
		......
		..###.
		......
		.....g

		want path:

		s***..
		###*..
		.***..
		.*###.
		.**...
		..***g
	*/

	grid := []string{
		"s.....",
		"###...",
		"......",
		"..###.",
		"......",
		".....g",
	}

	start := coord{0, 0}
	goal := coord{5, 5}

	// coords holds all traversible coordinates
	coords := make(map[coord]struct{})
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] != '#' {
				coords[coord{x, y}] = struct{}{}
			}
		}
	}

	neighbours := func(c coord) []coord {
		var ns []coord

		for _, n := range []coord{
			{c.x - 1, c.y},
			{c.x + 1, c.y},
			{c.x, c.y - 1},
			{c.x, c.y + 1},
		} {
			if _, ok := coords[n]; ok {
				ns = append(ns, n)
			}
		}

		return ns
	}

	cost := func(a, b coord) int {
		if a.x == 5 || b.x == 5 {
			return 10
		}
		return 1
	}

	heuristic := func(c coord) int {
		return manhattanDistance(c, goal)
	}

	gotPath, found := aStar(start, goal, neighbours, cost, heuristic)

	if got, want := found, true; got != want {
		t.Fatalf("found is %t, want %t", got, want)
	}

	wantPath := []coord{
		{0, 0}, {1, 0}, {2, 0}, {3, 0}, {3, 1},
		{3, 2}, {2, 2}, {1, 2}, {1, 3}, {1, 4},
		{2, 4}, {2, 5}, {3, 5}, {4, 5}, {5, 5},
	}

	if fmt.Sprint(gotPath) != fmt.Sprint(wantPath) {
		t.Errorf("got path:\n%v\n\nwant:\n%v", gotPath, wantPath)
	}
}

func TestManhattanDistance(t *testing.T) {
	for n, tc := range []struct {
		a, b     coord
		distance int
	}{
		{coord{0, 0}, coord{0, 0}, 0},

		{coord{0, 0}, coord{0, -1}, 1},
		{coord{0, 0}, coord{-1, -1}, 2},
		{coord{0, 0}, coord{-1, 0}, 1},
		{coord{0, 0}, coord{-1, 1}, 2},
		{coord{0, 0}, coord{0, 1}, 1},
		{coord{0, 0}, coord{1, 1}, 2},
		{coord{0, 0}, coord{1, 0}, 1},
		{coord{0, 0}, coord{1, -1}, 2},

		{coord{0, 0}, coord{0, -2}, 2},
		{coord{0, 0}, coord{-1, -2}, 3},
		{coord{0, 0}, coord{-2, -2}, 4},
		{coord{0, 0}, coord{-2, -1}, 3},
		{coord{0, 0}, coord{-2, 0}, 2},
		{coord{0, 0}, coord{-2, 1}, 3},
		{coord{0, 0}, coord{-2, 2}, 4},
		{coord{0, 0}, coord{-1, 2}, 3},
		{coord{0, 0}, coord{0, 2}, 2},
		{coord{0, 0}, coord{1, 2}, 3},
		{coord{0, 0}, coord{2, 2}, 4},
		{coord{0, 0}, coord{2, 1}, 3},
		{coord{0, 0}, coord{2, 0}, 2},
		{coord{0, 0}, coord{2, -1}, 3},
		{coord{0, 0}, coord{2, -2}, 4},
		{coord{0, 0}, coord{1, -2}, 3},
	} {
		if got, want := manhattanDistance(tc.a, tc.b), tc.distance; got != want {
			t.Errorf("[%d] manhattanDistance(%v, %v) = %d, want %d", n, tc.a, tc.b, got, want)
		}
	}
}
