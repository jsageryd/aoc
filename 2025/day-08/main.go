package main

import (
	"bufio"
	"cmp"
	"fmt"
	"maps"
	"math"
	"os"
	"slices"
)

func main() {
	var input []string

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		input = append(input, s.Text())
	}

	fmt.Printf("Part 1: %d\n", part1(input, 1000))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string, limit int) int {
	circuits, _ := find(input, limit)

	slices.SortFunc(circuits, func(a, b map[box]bool) int {
		return cmp.Compare(len(b), len(a))
	})

	m := 1

	for n := range 3 {
		m *= len(circuits[n])
	}

	return m
}

func part2(input []string) int {
	_, lastPair := find(input, -1)

	return int(lastPair[0].x) * int(lastPair[1].x)
}

func find(input []string, limit int) (circuits []map[box]bool, lastPair [2]box) {
	boxes := parse(input)
	pairs := findPairs(boxes)

	// Sort pairs descending by distance
	slices.SortFunc(pairs, func(a, b [2]box) int {
		return cmp.Compare(distance(a[0], a[1]), distance(b[0], b[1]))
	})

	for _, b := range boxes {
		circuits = append(circuits, map[box]bool{b: true})
	}

	if limit > -1 {
		pairs = pairs[:limit]
	}

	for _, pair := range pairs {
		idxA := slices.IndexFunc(circuits, func(circuit map[box]bool) bool {
			return circuit[pair[0]]
		})

		idxB := slices.IndexFunc(circuits, func(circuit map[box]bool) bool {
			return circuit[pair[1]]
		})

		// If both are already in the same circuit, skip
		if idxA == idxB {
			continue
		}

		// Merge the circuits
		maps.Insert(circuits[idxA], maps.All(circuits[idxB]))
		circuits = slices.Delete(circuits, idxB, idxB+1)

		lastPair = pair

		if len(circuits) == 1 {
			break
		}
	}

	return circuits, lastPair
}

// findPairs returns all possible box pairs.
func findPairs(boxes []box) [][2]box {
	var pairs [][2]box

	for n, a := range boxes {
		for _, b := range boxes[n+1:] {
			pairs = append(pairs, [2]box{a, b})
		}
	}

	return pairs
}

type box struct {
	x, y, z float64
}

func distance(a, b box) float64 {
	dx := a.x - b.x
	dy := a.y - b.y
	dz := a.z - b.z

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func parse(input []string) []box {
	var boxes []box

	for _, line := range input {
		var b box
		fmt.Sscanf(line, "%f,%f,%f", &b.x, &b.y, &b.z)
		boxes = append(boxes, b)
	}

	return boxes
}
