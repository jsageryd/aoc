package main

import (
	"fmt"
)

func main() {
	var lines []line

	for {
		var l line
		if _, err := fmt.Scanf(
			"%d,%d -> %d,%d",
			&l.from.x, &l.from.y, &l.to.x, &l.to.y,
		); err != nil {
			break
		}
		lines = append(lines, l)
	}

	fmt.Printf("Part 1: %d\n", countOverlaps(lines, false))
	fmt.Printf("Part 2: %d\n", countOverlaps(lines, true))
}

func countOverlaps(lines []line, countDiagonals bool) int {
	var overlaps int
	for _, n := range drawLines(lines, countDiagonals) {
		if n > 1 {
			overlaps++
		}
	}
	return overlaps
}

func drawLines(lines []line, drawDiagonals bool) map[coord]int {
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	grid := make(map[coord]int)

	for _, l := range lines {
		x1, y1, x2, y2 := l.from.x, l.from.y, l.to.x, l.to.y

		if !drawDiagonals && !(x1 == x2 || y1 == y2) {
			continue
		}

		steps := abs(x1-x2) + 1
		if steps == 1 {
			steps = abs(y1-y2) + 1
		}
		x, y := x1, y1
		for i := 0; i < steps; i++ {
			grid[coord{x, y}]++
			if x1 < x2 {
				x++
			} else if x1 > x2 {
				x--
			}
			if y1 < y2 {
				y++
			} else if y1 > y2 {
				y--
			}
		}
	}

	return grid
}

type line struct {
	from coord
	to   coord
}

type coord struct {
	x, y int
}
