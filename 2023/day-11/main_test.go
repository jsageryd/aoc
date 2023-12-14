package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	input := []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}

	if got, want := part1(input), 374; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestCombinations(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	k := 3

	wantCombs := [][]int{
		{1, 2, 3},
		{1, 2, 4},
		{1, 2, 5},
		{1, 2, 6},
		{1, 3, 4},
		{1, 3, 5},
		{1, 3, 6},
		{1, 4, 5},
		{1, 4, 6},
		{1, 5, 6},
		{2, 3, 4},
		{2, 3, 5},
		{2, 3, 6},
		{2, 4, 5},
		{2, 4, 6},
		{2, 5, 6},
		{3, 4, 5},
		{3, 4, 6},
		{3, 5, 6},
		{4, 5, 6},
	}

	var gotCombs [][]int

	combinations(s, k, func(comb []int) bool {
		c := make([]int, k)
		copy(c, comb)
		gotCombs = append(gotCombs, c)
		return true
	})

	if fmt.Sprint(gotCombs) != fmt.Sprint(wantCombs) {
		t.Errorf("got:\n%v\nwant:\n%v", gotCombs, wantCombs)
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
