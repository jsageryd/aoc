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

func traverseForest(forest []string, vx, vy int, f func(x, y int, c byte)) {
	var x, y int
	for _, row := range forest {
		f(x, y, row[x%len(row)])
		x, y = x+vx, y+vy
	}
}
