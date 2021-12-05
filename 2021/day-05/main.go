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

	fmt.Printf("Part 1: %d\n", countOverlaps(lines))
}

func countOverlaps(lines []line) int {
	var overlaps int
	for _, n := range drawLines(lines) {
		if n > 1 {
			overlaps++
		}
	}
	return overlaps
}

func drawLines(lines []line) map[coord]int {
	grid := make(map[coord]int)
	for _, l := range lines {
		if !(l.from.x == l.to.x || l.from.y == l.to.y) {
			continue
		}
		x1, y1, x2, y2 := l.from.x, l.from.y, l.to.x, l.to.y
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for y := y1; y <= y2; y++ {
			for x := x1; x <= x2; x++ {
				grid[coord{x, y}]++
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
