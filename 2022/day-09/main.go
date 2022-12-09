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

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	tailLocations := make(map[coord]struct{})

	var h, t coord

	for _, motion := range input {
		dist, _ := strconv.Atoi(motion[2:])
		for n := 0; n < dist; n++ {
			h.move(string(motion[0]))
			t.follow(h)
			tailLocations[t] = struct{}{}
		}
	}

	return len(tailLocations)
}

func part2(input []string) int {
	tailLocations := make(map[coord]struct{})

	knots := make([]coord, 10)

	for _, motion := range input {
		dist, _ := strconv.Atoi(motion[2:])
		for n := 0; n < dist; n++ {
			knots[0].move(string(motion[0]))
			for m := 1; m < len(knots); m++ {
				knots[m].follow(knots[m-1])
			}
			tailLocations[knots[9]] = struct{}{}
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
	for !c.near(other) {
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

func (c *coord) near(other coord) bool {
	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	return abs(c.x-other.x) < 2 && abs(c.y-other.y) < 2
}
