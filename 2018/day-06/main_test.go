package main

import (
	"fmt"
	"os"
	"testing"
)

func TestLargestNonInfiniteTerritory(t *testing.T) {
	coords := []coord{{1, 1}, {1, 6}, {8, 3}, {3, 4}, {5, 5}, {8, 9}}

	if got, want := largestNonInfiniteTerritory(coords), 17; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSizeOfTerritoryClosestToAllCoords(t *testing.T) {
	coords := []coord{{1, 1}, {1, 6}, {8, 3}, {3, 4}, {5, 5}, {8, 9}}

	if got, want := sizeOfTerritoryClosestToAllCoords(coords, 32), 16; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestBoundaries(t *testing.T) {
	for n, tc := range []struct {
		coords      []coord
		topLeft     coord
		bottomRight coord
	}{
		{
			coords:      []coord{{0, 0}},
			topLeft:     coord{0, 0},
			bottomRight: coord{0, 0},
		},
		{
			coords:      []coord{{0, 0}, {1, 1}},
			topLeft:     coord{0, 0},
			bottomRight: coord{1, 1},
		},
		{
			coords:      []coord{{0, 0}, {1, 1}},
			topLeft:     coord{0, 0},
			bottomRight: coord{1, 1},
		},
		{
			coords:      []coord{{0, 0}, {1, 1}, {2, 2}},
			topLeft:     coord{0, 0},
			bottomRight: coord{2, 2},
		},
		{
			coords:      []coord{{1, 1}, {2, 2}, {0, 0}},
			topLeft:     coord{0, 0},
			bottomRight: coord{2, 2},
		},
		{
			coords:      []coord{{5, 4}, {2, 6}, {1, 0}, {0, 1}},
			topLeft:     coord{0, 0},
			bottomRight: coord{5, 6},
		},
	} {
		topLeft, bottomRight := boundaries(tc.coords)

		if topLeft != tc.topLeft || bottomRight != tc.bottomRight {
			t.Errorf("[%d] boundaries(%v) = %v, %v; want %v, %v", n, tc.coords, topLeft, bottomRight, tc.topLeft, tc.bottomRight)
		}
	}
}

func TestClosest(t *testing.T) {
	for n, tc := range []struct {
		c       coord
		coords  []coord
		closest coord
		ok      bool
	}{
		{
			c:       coord{0, 0},
			coords:  []coord{{0, 0}},
			closest: coord{0, 0},
			ok:      true,
		},
		{
			c:       coord{0, 0},
			coords:  []coord{{1, 1}},
			closest: coord{1, 1},
			ok:      true,
		},
		{
			c:       coord{1, 0},
			coords:  []coord{{0, 0}, {2, 0}},
			closest: coord{0, 0},
			ok:      false,
		},
		{
			c:       coord{0, 0},
			coords:  []coord{{0, 0}, {0, 0}},
			closest: coord{0, 0},
			ok:      false,
		},
		{
			c:       coord{0, 0},
			coords:  []coord{{-1, -1}, {1, 1}},
			closest: coord{0, 0},
			ok:      false,
		},
		{
			c:       coord{0, -1},
			coords:  []coord{{-1, -1}, {1, 1}},
			closest: coord{-1, -1},
			ok:      true,
		},
		{
			c:       coord{0, 3},
			coords:  []coord{{0, 0}, {0, 4}, {4, 0}, {4, 4}},
			closest: coord{0, 4},
			ok:      true,
		},
		{
			c:       coord{5, 0},
			coords:  []coord{{1, 1}, {1, 6}, {8, 3}, {3, 4}, {5, 5}, {8, 9}},
			closest: coord{0, 0},
			ok:      false,
		},
		{
			c:       coord{5, 1},
			coords:  []coord{{1, 1}, {1, 6}, {8, 3}, {3, 4}, {5, 5}, {8, 9}},
			closest: coord{0, 0},
			ok:      false,
		},
		{
			c:       coord{0, 4},
			coords:  []coord{{1, 1}, {1, 6}, {8, 3}, {3, 4}, {5, 5}, {8, 9}},
			closest: coord{0, 0},
			ok:      false,
		},
		{
			c:       coord{1, 4},
			coords:  []coord{{1, 1}, {1, 6}, {8, 3}, {3, 4}, {5, 5}, {8, 9}},
			closest: coord{0, 0},
			ok:      false,
		},
	} {
		closest, ok := closest(tc.c, tc.coords)

		if closest != tc.closest || ok != tc.ok {
			t.Errorf("[%d] closest(%v, %v) = %v, %t; want %v, %t", n, tc.c, tc.coords, closest, ok, tc.closest, tc.ok)
		}
	}
}

func TestDistance(t *testing.T) {
	for n, tc := range []struct {
		c1, c2   coord
		distance int
	}{
		{c1: coord{0, 0}, c2: coord{0, 0}, distance: 0},
		{c1: coord{0, 0}, c2: coord{0, 1}, distance: 1},
		{c1: coord{0, 0}, c2: coord{1, 1}, distance: 2},
		{c1: coord{0, 1}, c2: coord{0, 0}, distance: 1},
		{c1: coord{1, 1}, c2: coord{0, 0}, distance: 2},
		{c1: coord{5, 5}, c2: coord{6, 6}, distance: 2},
		{c1: coord{1, 5}, c2: coord{5, 1}, distance: 8},
		{c1: coord{-1, -5}, c2: coord{-5, -1}, distance: 8},
	} {
		if got, want := distance(tc.c1, tc.c2), tc.distance; got != want {
			t.Errorf("[%d] distance(%v, %v) = %d, want %d", n, tc.c1, tc.c2, got, want)
		}
	}
}

func BenchmarkLargestNonInfiniteTerritory(b *testing.B) {
	f, _ := os.Open("input")
	defer f.Close()

	var input []coord

	for {
		var x, y int
		if _, err := fmt.Fscanf(f, "%d, %d", &x, &y); err != nil {
			break
		}
		input = append(input, coord{x, y})
	}

	for n := 0; n < b.N; n++ {
		largestNonInfiniteTerritory(input)
	}
}

func BenchmarkSizeOfTerritoryClosestToAllCoords(b *testing.B) {
	f, _ := os.Open("input")
	defer f.Close()

	var input []coord

	for {
		var x, y int
		if _, err := fmt.Fscanf(f, "%d, %d", &x, &y); err != nil {
			break
		}
		input = append(input, coord{x, y})
	}

	for n := 0; n < b.N; n++ {
		sizeOfTerritoryClosestToAllCoords(input, 10000)
	}
}
