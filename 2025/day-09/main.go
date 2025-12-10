package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input []string

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		input = append(input, s.Text())
	}

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input []string) int {
	var maxArea int

	for _, pair := range findPairs(parse(input)) {
		if a := area(pair[0], pair[1]); a > maxArea {
			maxArea = a
		}
	}

	return maxArea
}

type coord struct {
	x, y int
}

func area(a, b coord) int {
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	return (abs(a.x-b.x) + 1) * (abs(a.y-b.y) + 1)
}

// findPairs returns all possible tile pairs.
func findPairs(tiles []coord) [][2]coord {
	var pairs [][2]coord

	for n, a := range tiles {
		for _, b := range tiles[n+1:] {
			pairs = append(pairs, [2]coord{a, b})
		}
	}

	return pairs
}

func parse(input []string) []coord {
	var tiles []coord

	for _, line := range input {
		var tile coord
		fmt.Sscanf(line, "%d,%d", &tile.x, &tile.y)
		tiles = append(tiles, tile)
	}

	return tiles
}
