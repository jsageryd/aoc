package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Printf("Part 1: %d\n", part1(input))
}

func part1(input []string) int {
	cave := newCave(input)

	for cave.step() {
	}

	return cave.amountOfSand()
}

type coord struct {
	x, y int
}

type Cave struct {
	grid                 map[coord]byte
	sandSource           coord
	topLeft, bottomRight coord
	grain                coord
}

func newCave(input []string) *Cave {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	sandSource := coord{500, 0}

	cave := &Cave{
		grid:        map[coord]byte{sandSource: '+'},
		sandSource:  sandSource,
		topLeft:     sandSource,
		bottomRight: sandSource,
		grain:       sandSource,
	}

	updateBounds := func(wall coord) {
		if wall.x < cave.topLeft.x {
			cave.topLeft.x = wall.x
		}
		if wall.x > cave.bottomRight.x {
			cave.bottomRight.x = wall.x
		}
		if wall.y < cave.topLeft.y {
			cave.topLeft.y = wall.y
		}
		if wall.y > cave.bottomRight.y {
			cave.bottomRight.y = wall.y
		}
	}

	for _, line := range input {
		var last coord

		for n, cStr := range strings.Split(line, " -> ") {
			var cur coord

			fmt.Sscanf(cStr, "%d,%d", &cur.x, &cur.y)

			updateBounds(cur)

			if n != 0 {
				fromX, toX := min(last.x, cur.x), max(last.x, cur.x)
				fromY, toY := min(last.y, cur.y), max(last.y, cur.y)

				for y := fromY; y <= toY; y++ {
					for x := fromX; x <= toX; x++ {
						cave.grid[coord{x, y}] = '#'
					}
				}
			}

			last = cur
		}
	}

	return cave
}

// step moves forward one unit of time and returns true as long as sand flows.
func (c *Cave) step() bool {
	rest := true

	delete(c.grid, c.grain)

	for _, dxdy := range []coord{
		{0, 1}, {-1, 1}, {1, 1},
	} {
		dx, dy := dxdy.x, dxdy.y

		_, foundObstacle := c.grid[coord{c.grain.x + dx, c.grain.y + dy}]

		if !foundObstacle {
			c.grain.x += dx
			c.grain.y += dy
			rest = false
			break
		}
	}

	if c.grain.y > c.bottomRight.y {
		c.grain = c.sandSource
		return false
	}

	c.grid[c.grain] = 'o'

	if rest {
		c.grain = c.sandSource
	}

	return true
}

func (c *Cave) amountOfSand() int {
	var grains int

	for _, t := range c.grid {
		if t == 'o' {
			grains++
		}
	}

	return grains
}

func (c *Cave) String() string {
	var b strings.Builder

	for y := c.topLeft.y; y <= c.bottomRight.y; y++ {
		for x := c.topLeft.x; x <= c.bottomRight.x; x++ {
			switch t := c.grid[coord{x, y}]; t {
			case '#', 'o':
				b.WriteByte(t)
			default:
				if (coord{x, y}) == c.sandSource {
					b.WriteByte('+')
				} else {
					b.WriteByte('.')
				}
			}
		}
		if y != c.bottomRight.y {
			b.WriteByte('\n')
		}
	}

	return b.String()
}
