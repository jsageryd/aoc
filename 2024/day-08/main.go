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
	antennas := parse(input)

	antinodes := make(map[Coord]struct{})

	for _, locations := range antennas {
		for _, loc := range locations {
			for _, loc2 := range locations {
				if loc == loc2 {
					continue
				}

				dx := loc.x - loc2.x
				dy := loc.y - loc2.y

				antinode := Coord{loc.x + dx, loc.y + dy}

				if antinode.x >= 0 && antinode.x < len(input[0]) &&
					antinode.y >= 0 && antinode.y < len(input) {
					antinodes[antinode] = struct{}{}
				}
			}
		}
	}

	return len(antinodes)
}

func part2(input []string) int {
	antennas := parse(input)

	antinodes := make(map[Coord]struct{})

	for _, locations := range antennas {
		for _, loc := range locations {
			for _, loc2 := range locations {
				if loc == loc2 {
					continue
				}

				dx := loc.x - loc2.x
				dy := loc.y - loc2.y

				for x, y := loc.x, loc.y; 0 <= x && x < len(input[0]) && 0 <= y && y < len(input); x, y = x+dx, y+dy {
					antinodes[Coord{x, y}] = struct{}{}
				}
			}
		}
	}

	return len(antinodes)
}

type Coord struct {
	x, y int
}

func parse(input []string) (antennas map[byte][]Coord) {
	antennas = make(map[byte][]Coord)

	for y := range input {
		for x := range input[y] {
			if input[y][x] == '.' {
				continue
			}

			antennas[input[y][x]] = append(antennas[input[y][x]], Coord{x, y})
		}
	}

	return antennas
}
