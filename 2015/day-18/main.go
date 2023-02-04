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

	grid := parse(input, false)
	step(grid, 100, false)

	fmt.Printf("Part 1: %d\n", countOn(grid))

	grid = parse(input, true)
	step(grid, 100, true)

	fmt.Printf("Part 2: %d\n", countOn(grid))
}

type coord struct {
	x, y int
}

func parse(input []string, cornersStuck bool) map[coord]bool {
	grid := make(map[coord]bool)

	for y := range input {
		for x := range input[y] {
			grid[coord{x, y}] = input[y][x] == '#'
		}
	}

	if cornersStuck {
		tl, tr, bl, br := corners(grid)

		for _, c := range []coord{tl, tr, bl, br} {
			grid[c] = true
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

func step(grid map[coord]bool, steps int, cornersStuck bool) {
	toggle := make(map[coord]struct{})

	tl, tr, bl, br := corners(grid)

	for n := 0; n < steps; n++ {
		for c, on := range grid {
			if cornersStuck {
				switch c {
				case tl, tr, bl, br:
					continue
				}
			}

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

func corners(grid map[coord]bool) (topLeft, topRight, bottomLeft, bottomRight coord) {
	for c := range grid {
		if c.x > bottomRight.x {
			bottomRight.x = c.x
		}
		if c.y > bottomRight.y {
			bottomRight.y = c.y
		}
	}

	return coord{0, 0}, coord{bottomRight.x, 0}, coord{0, bottomRight.y}, bottomRight
}

func draw(grid map[coord]bool) string {
	_, _, _, max := corners(grid)

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
