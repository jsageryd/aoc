package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	s := newShip()
	navigate(s, input)
	return abs(s.x) + abs(s.y)
}

func part2(input []string) int {
	s := newShip()
	navigateWithWaypoint(s, input)
	return abs(s.x) + abs(s.y)
}

type ship struct {
	x, y    int
	heading byte
}

func newShip() *ship {
	return &ship{heading: 'E'}
}

func navigate(s *ship, instructions []string) {
	move := func(heading byte, distance int) {
		switch heading {
		case 'N':
			s.y -= distance
		case 'E':
			s.x += distance
		case 'S':
			s.y += distance
		case 'W':
			s.x -= distance
		}
	}

	turn := func(degrees int) {
		const nesw = "NESW"
		cur := strings.IndexByte(nesw, s.heading)
		for degrees < 0 {
			degrees += 360
		}
		s.heading = nesw[(cur+(degrees/90))%4]
	}

	for _, inst := range instructions {
		var action byte
		var value int

		fmt.Sscanf(inst, "%c%d", &action, &value)

		switch action {
		case 'N', 'E', 'S', 'W':
			move(action, value)
		case 'L':
			turn(-value)
		case 'R':
			turn(value)
		case 'F':
			move(s.heading, value)
		}
	}
}

func navigateWithWaypoint(s *ship, instructions []string) {
	waypoint := struct {
		dx, dy int
	}{10, -1}

	moveWaypoint := func(heading byte, distance int) {
		switch heading {
		case 'N':
			waypoint.dy -= distance
		case 'E':
			waypoint.dx += distance
		case 'S':
			waypoint.dy += distance
		case 'W':
			waypoint.dx -= distance
		}
	}

	rotateWaypoint := func(degrees int) {
		const cos90 = 0
		const sin90 = 1

		for degrees < 0 {
			degrees += 360
		}

		for i := 0; i < (degrees/90)%4; i++ {
			newX := waypoint.dx*cos90 - waypoint.dy*sin90
			newY := waypoint.dx*sin90 + waypoint.dy*cos90
			waypoint.dx, waypoint.dy = newX, newY
		}

		// https://en.wikipedia.org/wiki/Rotation_matrix#In_two_dimensions
	}

	for _, inst := range instructions {
		var action byte
		var value int

		fmt.Sscanf(inst, "%c%d", &action, &value)

		switch action {
		case 'N', 'E', 'S', 'W':
			moveWaypoint(action, value)
		case 'L':
			rotateWaypoint(-value)
		case 'R':
			rotateWaypoint(value)
		case 'F':
			s.x += value * waypoint.dx
			s.y += value * waypoint.dy
		}
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
