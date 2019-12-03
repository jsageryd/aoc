package main

import (
	"fmt"
	"testing"
)

func TestDrawWire(t *testing.T) {
	for n, tc := range []struct {
		path string
		want []coord
	}{
		{
			/*
			   ...........
			   ...........
			   ...........
			   ....+----+.
			   ....|....|.
			   ....|....|.
			   ....|....|.
			   .........|.
			   .o-------+.
			   ...........
			*/
			path: "R8,U5,L5,D3",
			want: []coord{
				{0, 0},

				// R8
				{1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {7, 0}, {8, 0},

				// U5
				{8, -1}, {8, -2}, {8, -3}, {8, -4}, {8, -5},

				// L5
				{7, -5}, {6, -5}, {5, -5}, {4, -5}, {3, -5},

				// D3
				{3, -4}, {3, -3}, {3, -2},
			},
		},
		{
			/*
				.+-----+...
				.|.....|...
				.|.....|...
				.|.....|...
				.|.----+...
				.|.........
				.|.........
				.o.........
				...........
			*/
			path: "U7,R6,D4,L4",
			want: []coord{
				{0, 0},

				// U7
				{0, -1}, {0, -2}, {0, -3}, {0, -4}, {0, -5}, {0, -6}, {0, -7},

				// R6
				{1, -7}, {2, -7}, {3, -7}, {4, -7}, {5, -7}, {6, -7},

				// D4
				{6, -6}, {6, -5}, {6, -4}, {6, -3},

				// L4
				{5, -3}, {4, -3}, {3, -3}, {2, -3},
			},
		},
	} {
		got := drawWire(tc.path)

		if fmt.Sprint(got) != fmt.Sprint(tc.want) {
			t.Errorf("[%d] got %v, want %v", n, got, tc.want)
		}
	}
}

func TestIntersections(t *testing.T) {
	for n, tc := range []struct {
		w1, w2 []coord
		want   []coord
	}{
		{
			w1: drawWire("R8,U5,L5,D3"),
			w2: drawWire("U7,R6,D4,L4"),
			want: []coord{
				{3, -3},
				{6, -5},
			},
		},
	} {
		got := intersections(tc.w1, tc.w2)

		if fmt.Sprint(got) != fmt.Sprint(tc.want) {
			t.Errorf("[%d] got %v, want %v", n, got, tc.want)
		}
	}
}

func TestDistance(t *testing.T) {
	for n, tc := range []struct {
		c1, c2   coord
		distance int
	}{
		{c1: coord{0, 0}, c2: coord{0, 0}, distance: 0},

		{c1: coord{0, 0}, c2: coord{1, 0}, distance: 1},
		{c1: coord{0, 0}, c2: coord{0, 1}, distance: 1},
		{c1: coord{1, 0}, c2: coord{0, 0}, distance: 1},
		{c1: coord{0, 1}, c2: coord{0, 0}, distance: 1},

		{c1: coord{0, 0}, c2: coord{-1, 0}, distance: 1},
		{c1: coord{0, 0}, c2: coord{0, -1}, distance: 1},
		{c1: coord{-1, 0}, c2: coord{0, 0}, distance: 1},
		{c1: coord{0, -1}, c2: coord{0, 0}, distance: 1},

		{c1: coord{0, 0}, c2: coord{1, 1}, distance: 2},
		{c1: coord{1, 1}, c2: coord{0, 0}, distance: 2},

		{c1: coord{-1, -1}, c2: coord{1, 1}, distance: 4},
	} {
		if got, want := distance(tc.c1, tc.c2), tc.distance; got != want {
			t.Errorf("[%d] distance(%v, %v) = %d, want %d", n, tc.c1, tc.c2, got, want)
		}
	}
}

func TestClosestIntersection(t *testing.T) {
	for n, tc := range []struct {
		w1, w2  []coord
		closest coord
	}{
		{
			w1:      drawWire("R75,D30,R83,U83,L12,D49,R71,U7,L72"),
			w2:      drawWire("U62,R66,U55,R34,D71,R55,D58,R83"),
			closest: coord{155, -4},
		},
		{
			w1:      drawWire("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"),
			w2:      drawWire("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"),
			closest: coord{124, -11},
		},
	} {
		if got, want := closestIntersection(tc.w1, tc.w2), tc.closest; got != want {
			t.Errorf("[%d] got %v (distance %d), want %v", n, got, distance(coord{}, got), want)
		}
	}
}

func TestClosestIntersectionByWire(t *testing.T) {
	for n, tc := range []struct {
		w1, w2               []coord
		closest              coord
		combinedWireDistance int
	}{
		{
			w1:                   drawWire("R75,D30,R83,U83,L12,D49,R71,U7,L72"),
			w2:                   drawWire("U62,R66,U55,R34,D71,R55,D58,R83"),
			closest:              coord{158, 12},
			combinedWireDistance: 604,
		},
		{
			w1:                   drawWire("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"),
			w2:                   drawWire("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"),
			closest:              coord{107, -47},
			combinedWireDistance: 410,
		},
	} {
		if got, want := closestIntersectionByWire(tc.w1, tc.w2), tc.closest; got != want {
			t.Errorf("[%d] got %v (combined wire distance %d), want %v (%d)", n, got, idxOf(got, tc.w1)+idxOf(got, tc.w2), want, tc.combinedWireDistance)
		}
	}
}
