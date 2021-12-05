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
			&l.x1, &l.y1, &l.x2, &l.y2,
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
		if !drawDiagonals && !(l.x1 == l.x2 || l.y1 == l.y2) {
			continue
		}
		steps := abs(l.x1-l.x2) + 1
		if steps == 1 {
			steps = abs(l.y1-l.y2) + 1
		}
		x, y := l.x1, l.y1
		for i := 0; i < steps; i++ {
			grid[coord{x, y}]++
			if l.x1 < l.x2 {
				x++
			} else if l.x1 > l.x2 {
				x--
			}
			if l.y1 < l.y2 {
				y++
			} else if l.y1 > l.y2 {
				y--
			}
		}
	}

	return grid
}

type line struct {
	x1, y1, x2, y2 int
}

type coord struct {
	x, y int
}
