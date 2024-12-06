package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	if len(input) == 0 {
		return 0
	}

	guard, grid := parse(input)

	seen := make(map[Coord]bool)

	walk(guard, grid, func(guard Guard) bool {
		seen[guard.loc] = true
		return true
	})

	return len(seen)
}

func part2(input []string) int {
	if len(input) == 0 {
		return 0
	}

	guard, grid := parse(input)

	var sum int

	for y := range input {
		for x := range input[y] {
			if grid[Coord{x, y}] != '.' {
				continue
			}

			grid[Coord{x, y}] = 'O'

			seen := make(map[Guard]bool)

			walk(guard, grid, func(guard Guard) bool {
				if seen[guard] {
					sum++
					return false
				}

				seen[guard] = true

				return true
			})

			grid[Coord{x, y}] = '.'
		}
	}

	return sum
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

// walk moves the guard through the grid until f returns false.
func walk(guard Guard, grid map[Coord]byte, f func(Guard) bool) {
	for grid[guard.loc] != 0 && f(guard) {
		const dirs = "^>v<"

		next := map[byte]Coord{
			'^': Coord{guard.loc.x, guard.loc.y - 1},
			'>': Coord{guard.loc.x + 1, guard.loc.y},
			'v': Coord{guard.loc.x, guard.loc.y + 1},
			'<': Coord{guard.loc.x - 1, guard.loc.y},
		}

		if block := grid[next[guard.dir]]; block != '.' && block != 0 {
			guard.dir = dirs[(strings.IndexByte(dirs, guard.dir)+1)%4]
		} else {
			guard.loc = next[guard.dir]
		}
	}
}
