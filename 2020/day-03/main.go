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

	fmt.Printf("Part 1: %d\n", treeCount(input, 3, 1))
	fmt.Printf("Part 2: %d\n", multipliedTreeCounts(input))
}

func treeCount(forest []string, vx, vy int) int {
	var trees int
	traverseForest(forest, vx, vy, func(x, y int, c byte) {
		if c == '#' {
			trees++
		}
	})
	return trees
}

func multipliedTreeCounts(forest []string) int {
	vs := []struct {
		x, y int
	}{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	prod := treeCount(forest, vs[0].x, vs[0].y)
	for i := 1; i < len(vs); i++ {
		prod *= treeCount(forest, vs[i].x, vs[i].y)
	}

	return prod
}

func traverseForest(forest []string, vx, vy int, f func(x, y int, c byte)) {
	maxX := len(forest[0])
	for x, y := 0, 0; y < len(forest); x, y = x+vx, y+vy {
		f(x, y, forest[y][x%maxX])
	}
}
