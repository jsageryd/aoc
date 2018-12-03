package main

import (
	"fmt"
)

func main() {
	var input []int

	for {
		var n int
		if _, err := fmt.Scanln(&n); err != nil {
			break
		}
		input = append(input, n)
	}

	inputCopy := make([]int, len(input))
	copy(inputCopy, input)

	fmt.Printf("Part 1: %d\n", mazeSteps(input))
	fmt.Printf("Part 2: %d\n", mazeStepsAlternate(inputCopy))
}

func mazeSteps(offsets []int) int {
	steps := 0

	for n := 0; n >= 0 && n < len(offsets); {
		steps++
		n, offsets[n] = n+offsets[n], offsets[n]+1
	}

	return steps
}

func mazeStepsAlternate(offsets []int) int {
	steps := 0

	for n := 0; n >= 0 && n < len(offsets); {
		steps++
		inc := 1
		if offsets[n] == 3 {
			inc = -1
		}
		n, offsets[n] = n+offsets[n], offsets[n]+inc
	}

	return steps
}
