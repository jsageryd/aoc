package main

import "testing"

func TestTreeCount(t *testing.T) {
	forest := []string{
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

	if got, want := treeCount(forest, 3, 1), 7; got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
