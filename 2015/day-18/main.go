package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	grid := parse(input)
	step(grid, 100)

	fmt.Printf("Part 1: %d\n", countOn(grid))
}

type coord struct {
	x, y int
}

func parse(input []string) map[coord]bool {
	grid := make(map[coord]bool)

	for y := range input {
		for x := range input[y] {
			grid[coord{x, y}] = input[y][x] == '#'
		}
	}

	return grid
}

func countOn(grid map[coord]bool) int {
	var count int

	for _, on := range grid {
		if on {
			count++
		}
	}

	return count
}

func step(grid map[coord]bool, steps int) {
	toggle := make(map[coord]struct{})

	for n := 0; n < steps; n++ {
		for c, on := range grid {
			var neighboursOn int

			for _, dc := range []coord{
				{-1, -1}, {0, -1}, {1, -1},
				{-1, 0}, {1, 0},
				{-1, 1}, {0, 1}, {1, 1},
			} {
				if grid[coord{c.x + dc.x, c.y + dc.y}] {
					neighboursOn++
				}
			}

			if on {
				if neighboursOn != 2 && neighboursOn != 3 {
					toggle[c] = struct{}{}
				}
			} else {
				if neighboursOn == 3 {
					toggle[c] = struct{}{}
				}
			}
		}

		for c := range toggle {
			grid[c] = !grid[c]
			delete(toggle, c)
		}
	}
}

func draw(grid map[coord]bool) string {
	var max coord

	for c := range grid {
		if c.x > max.x {
			max.x = c.x
		}
		if c.y > max.y {
			max.y = c.y
		}
	}

	var buf bytes.Buffer

	for y := 0; y <= max.y; y++ {
		for x := 0; x <= max.x; x++ {
			if grid[coord{x, y}] {
				buf.WriteByte('#')
			} else {
				buf.WriteByte('.')
			}
		}
		if y != max.y {
			buf.WriteByte('\n')
		}
	}

	return buf.String()
}
