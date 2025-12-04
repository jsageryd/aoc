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
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	var count int

	grid := parse(input)

	for c := range grid {
		var nCount int

		for _, n := range c.neighbours() {
			if _, ok := grid[n]; ok {
				nCount++
			}
		}

		if nCount < 4 {
			count++
		}
	}

	return count
}

func part2(input []string) int {
	var count int

	grid := parse(input)

	var lastLen int
	var toRemove []coord

	for len(grid) != lastLen {
		for c := range grid {
			var nCount int

			for _, n := range c.neighbours() {
				if _, ok := grid[n]; ok {
					nCount++
				}
			}

			if nCount < 4 {
				toRemove = append(toRemove, c)
				count++
			}
		}

		lastLen = len(grid)

		for _, c := range toRemove {
			delete(grid, c)
		}
	}

	return count
}

type coord struct {
	x, y int
}

func (c coord) neighbours() []coord {
	return []coord{
		{c.x + 1, c.y},
		{c.x + 1, c.y - 1},
		{c.x, c.y - 1},
		{c.x - 1, c.y - 1},
		{c.x - 1, c.y},
		{c.x - 1, c.y + 1},
		{c.x, c.y + 1},
		{c.x + 1, c.y + 1},
	}
}

func parse(input []string) map[coord]struct{} {
	grid := make(map[coord]struct{})

	for y := range input {
		for x := range input[y] {
			if input[y][x] == '@' {
				grid[coord{x, y}] = struct{}{}
			}
		}
	}

	return grid
}
