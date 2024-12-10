package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
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
	var sum int

	grid := parse(input)

	for c, v := range grid {
		if v == 0 {
			seen := make(map[Coord]bool)

			trails := search(grid, c)

			for _, t := range trails {
				seen[t[len(t)-1]] = true
			}

			sum += len(seen)
		}
	}

	return sum
}

func part2(input []string) int {
	var sum int

	grid := parse(input)

	for c, v := range grid {
		if v == 0 {
			trails := search(grid, c)

			sum += len(trails)
		}
	}

	return sum
}

func search(grid map[Coord]int, trailhead Coord) (trails [][]Coord) {
	var rec func(trail []Coord)

	rec = func(trail []Coord) {
		tail := trail[len(trail)-1]

		if grid[tail] == 9 {
			trails = append(trails, trail)
			return
		}

		for _, neighbour := range []Coord{
			Coord{tail.x + 1, tail.y},
			Coord{tail.x, tail.y + 1},
			Coord{tail.x - 1, tail.y},
			Coord{tail.x, tail.y - 1},
		} {
			if grid[neighbour] == grid[tail]+1 {
				rec(slices.Clone(append(trail, neighbour)))
			}
		}
	}

	rec([]Coord{trailhead})

	return trails
}

type Coord struct {
	x, y int
}

func parse(input []string) map[Coord]int {
	grid := make(map[Coord]int)

	for y := range input {
		for x := range input[y] {
			v, _ := strconv.Atoi(string(input[y][x]))
			grid[Coord{x, y}] = v
		}
	}

	return grid
}
