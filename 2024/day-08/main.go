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
	antennas := make(map[byte][]Coord)

	for y := range input {
		for x := range input[y] {
			if input[y][x] == '.' {
				continue
			}

			antennas[input[y][x]] = append(antennas[input[y][x]], Coord{x, y})
		}
	}

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

type Coord struct {
	x, y int
}
