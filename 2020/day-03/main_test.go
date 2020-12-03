package main

import "testing"

var forest = []string{
	"..##.......",
	"#...#...#..",
	".#....#..#.",
	"..#.#...#.#",
	".#...##..#.",
	"..#.##.....",
	".#.#.#....#",
	".#........#",
	"#.##...#...",
	"#...##....#",
	".#..#...#.#",
}

func TestTreeCount(t *testing.T) {
	for n, tc := range []struct {
		vx, vy int
		trees  int
	}{
		0: {vx: 1, vy: 1, trees: 2},
		1: {vx: 3, vy: 1, trees: 7},
		2: {vx: 5, vy: 1, trees: 3},
		3: {vx: 7, vy: 1, trees: 4},
		4: {vx: 1, vy: 2, trees: 2},
	} {
		if got, want := treeCount(forest, tc.vx, tc.vy), tc.trees; got != want {
			t.Errorf("[%d] got %d, want %d", n, got, want)
		}
	}
}

func TestMultipliedTreeCounts(t *testing.T) {
	if got, want := multipliedTreeCounts(forest), 336; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
