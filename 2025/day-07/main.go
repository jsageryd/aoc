package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	var splits int

	grid := parse(input)

	var beams []coord

	for c, v := range grid {
		if v == 'S' {
			beams = append(beams, c)
			break
		}
	}

	for len(beams) > 0 {
		for n := len(beams) - 1; n >= 0; n-- {
			next := coord{beams[n].x, beams[n].y + 1}

			v, ok := grid[next]
			if !ok || v == '|' {
				beams = slices.Delete(beams, n, n+1)
				continue
			}

			switch v {
			case '^':
				splits++

				beams = slices.Delete(beams, n, n+1)

				for _, neighbour := range []coord{
					coord{next.x - 1, next.y},
					coord{next.x + 1, next.y},
				} {
					if v, ok := grid[neighbour]; ok && v != '|' {
						grid[neighbour] = '|'
						beams = append(beams, neighbour)
					}
				}
			default:
				grid[next] = '|'
				beams[n] = next
			}
		}
	}

	return splits
}

func part2(input []string) int {
	grid := parse(input)

	var start coord

	for c, v := range grid {
		if v == 'S' {
			start = c
			break
		}
	}

	mem := make(map[coord]int)

	var timelines func(next coord) int

	timelines = func(next coord) int {
		if c, ok := mem[next]; ok {
			return c
		}

		for {
			next = coord{next.x, next.y + 1}

			v, ok := grid[next]
			if !ok {
				return 1
			}

			if v == '^' {
				left := coord{next.x - 1, next.y}
				right := coord{next.x + 1, next.y}

				mem[left] = timelines(left)
				mem[right] = timelines(right)

				return mem[left] + mem[right]
			}
		}
	}

	return timelines(start)
}

type coord struct {
	x, y int
}

func parse(input []string) map[coord]byte {
	grid := make(map[coord]byte)

	for y := range input {
		for x := range input[y] {
			grid[coord{x, y}] = input[y][x]
		}
	}

	return grid
}

func gridString(grid map[coord]byte, beams []coord) string {
	var maxX, maxY int

	for c := range grid {
		if c.x > maxX {
			maxX = c.x
		}
		if c.y > maxY {
			maxY = c.y
		}
	}

	w, h := maxX+2, maxY+1

	b := make([]byte, h*w-1)

	for y := range h {
		for x := range w - 1 {
			b[y*w+x] = grid[coord{x, y}]
			for _, beam := range beams {
				if beam == (coord{x, y}) {
					b[y*w+x] = '+'
					break
				}
			}
		}
		if y < h-1 {
			b[y*w+w-1] = '\n'
		}
	}

	return string(b)
}
