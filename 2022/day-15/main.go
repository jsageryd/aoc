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

	fmt.Printf("Part 1: %d\n", part1(input, 2000000))
}

func part1(input []string, row int) int {
	m := parseInput(input)

	noBeacon := make(map[coord]struct{})

	md := manhattanDistance

	for sensor, beacon := range m {
		for x := sensor.x; md(sensor, coord{x, row}) <= md(sensor, beacon); x-- {
			noBeacon[coord{x, row}] = struct{}{}
		}

		for x := sensor.x + 1; md(sensor, coord{x, row}) <= md(sensor, beacon); x++ {
			noBeacon[coord{x, row}] = struct{}{}
		}

		delete(noBeacon, beacon)
	}

	return len(noBeacon)
}

type coord struct {
	x, y int
}

// parseInput parses the input into a map, sensor -> beacon.
func parseInput(input []string) map[coord]coord {
	m := make(map[coord]coord)

	for _, line := range input {
		var sensor, beacon coord
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensor.x, &sensor.y, &beacon.x, &beacon.y)
		m[sensor] = beacon
	}

	return m
}

// manhattanDistance returns the Manhattan distance between the given
// coordinates.
//
// Manhattan distances from center point:
//
//	43234
//	32123
//	21012
//	32123
//	43234
//
// https://en.wikipedia.org/wiki/Manhattan_distance
func manhattanDistance(a, b coord) int {
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	return abs(a.x-b.x) + abs(a.y-b.y)
}
