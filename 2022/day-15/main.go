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
	fmt.Printf("Part 2: %d\n", part2(input, 4000000))
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

func part2(input []string, max int) int {
	m := parseInput(input)

	// loopSensor loops around the sensor at the given radius, calling f for each
	// coordinate it visits, as long as f returns true.
	loopSensor := func(sensor coord, radius int, f func(c coord) bool) {
		// Start at 12 o'clock
		c := coord{
			x: sensor.x,
			y: sensor.y - radius - 1,
		}
		if !f(c) {
			return
		}

		// Counter-clockwise top-left quadrant
		for c.y < sensor.y {
			c.x--
			c.y++
			if !f(c) {
				return
			}
		}

		// Bottom-left quadrant
		for c.y < sensor.y+radius+1 {
			c.x++
			c.y++
			if !f(c) {
				return
			}
		}

		// Bottom-right quadrant
		for c.y > sensor.y {
			c.x++
			c.y--
			if !f(c) {
				return
			}
		}

		// Top-right quadrant
		for c.y > sensor.y-radius {
			c.x--
			c.y--
			if !f(c) {
				return
			}
		}
	}

	var missingBeacon coord

	md := manhattanDistance

	for sensor, beacon := range m {
		loopSensor(sensor, md(sensor, beacon), func(c coord) bool {
			// skip this coordinate if it's outside of the given constraints.
			if c.x < 0 || c.x > max || c.y < 0 || c.y > max {
				return true
			}

			for otherSensor, otherBeacon := range m {
				if otherSensor != sensor {
					if md(otherSensor, otherBeacon) >= md(otherSensor, c) {
						// This coordinate is within the radius of another sensor, so the
						// missing beacon cannot be here. Continue checking the next
						// coordinate.
						return true
					}
				}
			}

			// c was not within the radius of any other sensor, so this must be the
			// coordinate of the missing beacon.
			missingBeacon = c
			return false
		})
	}

	return missingBeacon.x*4000000 + missingBeacon.y
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
