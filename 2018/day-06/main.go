package main

import (
	"fmt"
	"sort"
)

func main() {
	var input []coord

	for {
		var x, y int
		if _, err := fmt.Scanf("%d, %d", &x, &y); err != nil {
			break
		}
		input = append(input, coord{x, y})
	}

	fmt.Printf("Part 1: %d\n", largestNonInfiniteTerritory(input))
	fmt.Printf("Part 2: %d\n", sizeOfTerritoryClosestToAllCoords(input, 10000))
}

type coord struct {
	x, y int
}

func largestNonInfiniteTerritory(coords []coord) int {
	territory := map[coord]int{}

	badCoords := map[coord]struct{}{}

	topLeft, bottomRight := boundaries(coords)

	for y := topLeft.y; y <= bottomRight.y; y++ {
		for x := topLeft.x; x <= bottomRight.x; x++ {
			if c, ok := closest(coord{x, y}, coords); ok {
				if y == topLeft.y || y == bottomRight.y || x == topLeft.x || x == bottomRight.x {
					badCoords[c] = struct{}{}
					continue
				}
				if _, ok := badCoords[c]; !ok {
					territory[c]++
				}
			}
		}
	}

	largestTerritory := 0

	for _, v := range territory {
		if v > largestTerritory {
			largestTerritory = v
		}
	}

	return largestTerritory
}

func sizeOfTerritoryClosestToAllCoords(coords []coord, maxSize int) int {
	territorySize := 0

	topLeft, bottomRight := boundaries(coords)

	for y := topLeft.y; y <= bottomRight.y; y++ {
		for x := topLeft.x; x <= bottomRight.x; x++ {
			totalDist := 0
			for n := range coords {
				totalDist += distance(coord{x, y}, coords[n])
			}
			if totalDist < maxSize {
				territorySize++
			}
		}
	}

	return territorySize
}

func boundaries(coords []coord) (topLeft, bottomRight coord) {
	topLeft = coords[0]
	bottomRight = coords[0]
	for _, c := range coords {
		if c.x < topLeft.x {
			topLeft.x = c.x
		}
		if c.y < topLeft.y {
			topLeft.y = c.y
		}
		if c.x > bottomRight.x {
			bottomRight.x = c.x
		}
		if c.y > bottomRight.y {
			bottomRight.y = c.y
		}
	}
	return topLeft, bottomRight
}

func closest(c coord, coords []coord) (closest coord, ok bool) {
	if len(coords) == 1 {
		return coords[0], true
	}
	sort.Slice(coords, func(i, j int) bool {
		return distance(c, coords[i]) < distance(c, coords[j])
	})
	if distance(c, coords[0]) == distance(c, coords[1]) {
		return coord{}, false
	}

	return coords[0], true
}

func distance(c1, c2 coord) int {
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	return abs(c1.x-c2.x) + abs(c1.y-c2.y)
}
