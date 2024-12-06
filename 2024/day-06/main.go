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
	if len(input) == 0 {
		return 0
	}

	guard, grid := parse(input)

	topLeft := Coord{0, 0}
	bottomRight := Coord{len(input[0]) - 1, len(input) - 1}

	seen := make(map[Coord]bool)

	for withinBounds(guard.loc, topLeft, bottomRight) {
		seen[guard.loc] = true

		switch guard.dir {
		case '^':
			if grid[Coord{guard.loc.x, guard.loc.y - 1}] == '#' {
				guard.dir = '>'
			} else {
				guard.loc.y--
			}
		case '<':
			if grid[Coord{guard.loc.x - 1, guard.loc.y}] == '#' {
				guard.dir = '^'
			} else {
				guard.loc.x--
			}
		case 'v':
			if grid[Coord{guard.loc.x, guard.loc.y + 1}] == '#' {
				guard.dir = '<'
			} else {
				guard.loc.y++
			}
		case '>':
			if grid[Coord{guard.loc.x + 1, guard.loc.y}] == '#' {
				guard.dir = 'v'
			} else {
				guard.loc.x++
			}
		}
	}

	return len(seen)
}

type Guard struct {
	dir byte
	loc Coord
}

type Coord struct {
	x, y int
}

func parse(input []string) (Guard, map[Coord]byte) {
	var guard Guard
	grid := make(map[Coord]byte)

	for y := range input {
		for x := range input[y] {
			if input[y][x] == '^' {
				guard = Guard{dir: input[y][x], loc: Coord{x, y}}
				grid[Coord{x, y}] = '.'
			} else {
				grid[Coord{x, y}] = input[y][x]
			}
		}
	}

	return guard, grid
}

func withinBounds(c Coord, topLeft, bottomRight Coord) bool {
	return c.x >= topLeft.x && c.x <= bottomRight.x &&
		c.y >= topLeft.y && c.y <= bottomRight.y
}
