package main

import (
	"fmt"
	"testing"
)

var input = []string{
	"Sensor at x=2, y=18: closest beacon is at x=-2, y=15",
	"Sensor at x=9, y=16: closest beacon is at x=10, y=16",
	"Sensor at x=13, y=2: closest beacon is at x=15, y=3",
	"Sensor at x=12, y=14: closest beacon is at x=10, y=16",
	"Sensor at x=10, y=20: closest beacon is at x=10, y=16",
	"Sensor at x=14, y=17: closest beacon is at x=10, y=16",
	"Sensor at x=8, y=7: closest beacon is at x=2, y=10",
	"Sensor at x=2, y=0: closest beacon is at x=2, y=10",
	"Sensor at x=0, y=11: closest beacon is at x=2, y=10",
	"Sensor at x=20, y=14: closest beacon is at x=25, y=17",
	"Sensor at x=17, y=20: closest beacon is at x=21, y=22",
	"Sensor at x=16, y=7: closest beacon is at x=15, y=3",
	"Sensor at x=14, y=3: closest beacon is at x=15, y=3",
	"Sensor at x=20, y=1: closest beacon is at x=15, y=3",
}

func TestPart1(t *testing.T) {
	if got, want := part1(input, 10), 26; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	if got, want := part2(input, 20), 56000011; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestParseInput(t *testing.T) {
	gotMap := parseInput(input)

	wantMap := map[coord]coord{
		{x: 2, y: 18}:  {x: -2, y: 15},
		{x: 9, y: 16}:  {x: 10, y: 16},
		{x: 13, y: 2}:  {x: 15, y: 3},
		{x: 12, y: 14}: {x: 10, y: 16},
		{x: 10, y: 20}: {x: 10, y: 16},
		{x: 14, y: 17}: {x: 10, y: 16},
		{x: 8, y: 7}:   {x: 2, y: 10},
		{x: 2, y: 0}:   {x: 2, y: 10},
		{x: 0, y: 11}:  {x: 2, y: 10},
		{x: 20, y: 14}: {x: 25, y: 17},
		{x: 17, y: 20}: {x: 21, y: 22},
		{x: 16, y: 7}:  {x: 15, y: 3},
		{x: 14, y: 3}:  {x: 15, y: 3},
		{x: 20, y: 1}:  {x: 15, y: 3},
	}

	if fmt.Sprint(gotMap) != fmt.Sprint(wantMap) {
		t.Errorf("got %v, want %v", gotMap, wantMap)
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
