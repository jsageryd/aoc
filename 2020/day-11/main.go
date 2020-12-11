package main

import (
	"bufio"
	"fmt"
	"os"
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
	layout := makeLayoutMap(input)

	for step(layout) {
	}

	var count int

	for _, v := range layout {
		if v == '#' {
			count++
		}
	}

	return count
}

func step(layout map[coord]byte) bool {
	var modified bool
	current := make(map[coord]byte, len(layout))

	for k, v := range layout {
		current[k] = v
	}

	adjacent := func(c coord) []coord {
		return []coord{
			{c.x - 1, c.y - 1},
			{c.x, c.y - 1},
			{c.x + 1, c.y - 1},
			{c.x - 1, c.y},
			{c.x + 1, c.y},
			{c.x - 1, c.y + 1},
			{c.x, c.y + 1},
			{c.x + 1, c.y + 1},
		}
	}

	adjacentOccupiedSeats := func(c coord) int {
		var count int
		for _, adj := range adjacent(c) {
			if current[adj] == '#' {
				count++
			}
		}
		return count
	}

	for k, v := range current {
		switch v {
		case 'L':
			if adjacentOccupiedSeats(k) == 0 {
				layout[k] = '#'
				modified = true
			}
		case '#':
			if adjacentOccupiedSeats(k) >= 4 {
				layout[k] = 'L'
				modified = true
			}
		}
	}

	return modified
}

type coord struct {
	x, y int
}

func makeLayoutMap(input []string) map[coord]byte {
	m := make(map[coord]byte)
	for y := range input {
		for x := range []byte(input[y]) {
			m[coord{x, y}] = input[y][x]
		}
	}
	return m
}
