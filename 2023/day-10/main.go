package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	m := parse(input)

	s := start(m)

	seen := map[coord]bool{s: true}

	ab := neighbours(m, s)
	a, b := ab[0], ab[1]

	var distance int

	for {
		distance++

		seen[a] = true
		seen[b] = true

		var okA, okB bool

		a, okA = next(m, a, seen)
		b, okB = next(m, b, seen)

		if !okA || !okB {
			break
		}
	}

	return distance
}

func part2(input []string) int {
	m := parse(input)

	s := start(m)

	seen := map[coord]bool{s: true}

	ab := neighbours(m, s)
	a, ok := ab[0], true

	m[s] = findType(m, s)

	for ok {
		seen[a] = true
		a, ok = next(m, a, seen)
	}

	var area int

	flip := map[byte]byte{
		'F': 'J',
		'J': 'F',
		'7': 'L',
		'L': '7',
	}

	var edgeStart byte

	for y := range input {
		var in bool
		var edge bool

		for x := range input[y] {
			c := coord{x, y}

			if seen[c] {
				switch m[c] {
				case '|':
					in = !in
				case 'F', '7', 'J', 'L':
					if !edge {
						edgeStart = m[c]
						in = !in
					} else {
						if m[c] != flip[edgeStart] {
							in = !in
						}
						edgeStart = 0
					}
					edge = !edge
				}
			} else {
				if in {
					area++
				}
			}
		}
	}

	return area
}

type coord struct {
	x, y int
}

func parse(input []string) map[coord]byte {
	m := make(map[coord]byte)

	for y := range input {
		for x := range input[y] {
			m[coord{x, y}] = input[y][x]
		}
	}

	return m
}

func start(m map[coord]byte) coord {
	for c, b := range m {
		if b == 'S' {
			return c
		}
	}
	return coord{}
}

func neighbours(m map[coord]byte, c coord) []coord {
	var (
		north = coord{c.x, c.y - 1}
		west  = coord{c.x - 1, c.y}
		south = coord{c.x, c.y + 1}
		east  = coord{c.x + 1, c.y}
	)

	switch m[c] {
	case '|':
		return []coord{north, south}
	case '-':
		return []coord{west, east}
	case 'L':
		return []coord{north, east}
	case 'J':
		return []coord{north, west}
	case '7':
		return []coord{south, west}
	case 'F':
		return []coord{south, east}
	case '.':
		return nil
	case 'S':
		var ns []coord

		if strings.IndexByte("|7F", m[north]) > -1 {
			ns = append(ns, north)
		}

		if strings.IndexByte("-LF", m[west]) > -1 {
			ns = append(ns, west)
		}

		if strings.IndexByte("|LJ", m[south]) > -1 {
			ns = append(ns, south)
		}

		if strings.IndexByte("-J7", m[east]) > -1 {
			ns = append(ns, east)
		}

		return ns
	default:
		return nil
	}
}

func next(m map[coord]byte, c coord, seen map[coord]bool) (n coord, ok bool) {
	for _, nc := range neighbours(m, c) {
		if !seen[nc] {
			return nc, true
		}
	}
	return coord{}, false
}

func findType(m map[coord]byte, c coord) byte {
	var (
		north = coord{c.x, c.y - 1}
		west  = coord{c.x - 1, c.y}
		south = coord{c.x, c.y + 1}
		east  = coord{c.x + 1, c.y}
	)

	var dirs []byte

	for _, c := range neighbours(m, c) {
		switch c {
		case north:
			dirs = append(dirs, 'n')
		case west:
			dirs = append(dirs, 'w')
		case south:
			dirs = append(dirs, 's')
		case east:
			dirs = append(dirs, 'e')
		}
	}

	slices.Sort(dirs)

	switch string(dirs) {
	case "en":
		return 'L'
	case "es":
		return '7'
	case "ew":
		return '-'
	case "ns":
		return '|'
	case "nw":
		return 'F'
	case "sw":
		return 'J'
	default:
		return 0
	}
}
