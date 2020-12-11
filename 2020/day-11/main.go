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

	fmt.Printf("Part 1: %d\n", run(input, step))
	fmt.Printf("Part 2: %d\n", run(input, stepV2))
}

func run(input []string, stepFunc func(layout map[coord]byte) bool) int {
	layout := makeLayoutMap(input)

	for stepFunc(layout) {
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

func stepV2(layout map[coord]byte) bool {
	var modified bool
	current := make(map[coord]byte, len(layout))

	for k, v := range layout {
		current[k] = v
	}

	walk := func(c coord, direction int) coord {
		switch direction {
		case 0: // north
			return coord{c.x, c.y - 1}
		case 1: // north-west
			return coord{c.x - 1, c.y - 1}
		case 2: // west
			return coord{c.x - 1, c.y}
		case 3: // south-west
			return coord{c.x - 1, c.y + 1}
		case 4: // south
			return coord{c.x, c.y + 1}
		case 5: // south-east
			return coord{c.x + 1, c.y + 1}
		case 6: // east
			return coord{c.x + 1, c.y}
		case 7: // north-east
			return coord{c.x + 1, c.y - 1}
		default:
			panic("unknown direction")
		}
	}

	occupiedSeatsInAnyDirection := func(c coord) int {
		var count int
		for d := 0; d < 8; d++ {
			cc := c
			for {
				cc = walk(cc, d)
				v, ok := current[cc]
				if !ok || v == 'L' {
					break
				}
				if v == '#' {
					count++
					break
				}
			}
		}
		return count
	}

	for k, v := range current {
		switch v {
		case 'L':
			if occupiedSeatsInAnyDirection(k) == 0 {
				layout[k] = '#'
				modified = true
			}
		case '#':
			if occupiedSeatsInAnyDirection(k) >= 5 {
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
