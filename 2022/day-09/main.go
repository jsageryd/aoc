package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	fmt.Printf("Part 1: %d\n", move(input, 2))
	fmt.Printf("Part 2: %d\n", move(input, 10))
}

func move(input []string, knotCount int) int {
	tailLocations := make(map[coord]struct{})

	knots := make([]coord, knotCount)

	for _, motion := range input {
		dist, _ := strconv.Atoi(motion[2:])
		for n := 0; n < dist; n++ {
			knots[0].move(string(motion[0]))
			for m := 1; m < len(knots); m++ {
				knots[m].follow(knots[m-1])
			}
			tailLocations[knots[len(knots)-1]] = struct{}{}
		}
	}

	return len(tailLocations)
}

type coord struct {
	x, y int
}

func (c *coord) move(direction string) {
	switch direction {
	case "L":
		c.x--
	case "R":
		c.x++
	case "U":
		c.y--
	case "D":
		c.y++
	}
}

func (c *coord) follow(other coord) {
	for chebyshevDistance(*c, other) > 1 {
		switch {
		case c.x < other.x:
			c.x++
		case c.x > other.x:
			c.x--
		}

		switch {
		case c.y < other.y:
			c.y++
		case c.y > other.y:
			c.y--
		}
	}
}

// chebyshevDistance returns the Chebyshev distance between the given
// coordinates.
//
// Chebyshev distances from center point:
//
//	22222
//	21112
//	21012
//	21112
//	22222
//
// https://en.wikipedia.org/wiki/Chebyshev_distance
func chebyshevDistance(a, b coord) int {
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	return max(abs(b.x-a.x), abs(b.y-a.y))
}
