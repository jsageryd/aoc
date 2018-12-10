package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var input []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	stars := parseInput(input)
	steps := findConstellation(stars)

	fmt.Println("Part 1:")
	printConstellation(stars)
	fmt.Println()

	fmt.Printf("Part 2: %d\n", steps)
}

func parseInput(input []string) []*star {
	var stars []*star

	for _, line := range input {
		var x, y, vx, vy int
		if _, err := fmt.Sscanf(
			line,
			"position=<%d, %d> velocity=<%d, %d>", &x, &y, &vx, &vy,
		); err != nil {
			break
		}
		stars = append(stars, &star{x, y, vx, vy})
	}

	return stars
}

type star struct {
	x, y   int
	vx, vy int
}

// findConstellation moves the stars until it arrives at the correct
// constellation and returns the number of steps needed to get there.
func findConstellation(stars []*star) (steps int) {
	area := totalArea(stars)
	var lastArea int
	for {
		move(stars, 1)
		lastArea = area
		area = totalArea(stars)
		if lastArea < area {
			break
		}
		steps++
	}
	move(stars, -1)
	return steps
}

func move(stars []*star, steps int) {
	for _, s := range stars {
		s.x += s.vx * steps
		s.y += s.vy * steps
	}
}

func totalArea(stars []*star) int {
	if len(stars) == 0 {
		return 0
	}
	x1 := stars[0].x
	y1 := stars[0].y
	x2, y2 := x1, y1
	for _, s := range stars {
		if s.x < x1 {
			x1 = s.x
		}
		if s.x > x2 {
			x2 = s.x
		}
		if s.y < y1 {
			y1 = s.y
		}
		if s.y > y2 {
			y2 = s.y
		}
	}
	return (x2 - x1) * (y2 - y1)
}

func printConstellation(stars []*star) {
	sort.Slice(stars, func(i, j int) bool {
		if stars[i].y == stars[j].y {
			return stars[i].x < stars[j].x
		}
		return stars[i].y < stars[j].y
	})

	x1, y1 := stars[0].x, stars[0].y
	x2, y2 := stars[len(stars)-1].x, stars[len(stars)-1].y

	starGrid := make([][]bool, y2-y1+1)
	for y := range starGrid {
		starGrid[y] = make([]bool, x2-x1+1)
	}
	for _, s := range stars {
		starGrid[s.y-y1][s.x-x1] = true
	}

	for y := range starGrid {
		for x := range starGrid[y] {
			if starGrid[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
