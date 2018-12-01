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

	fmt.Printf("Part 1: %d\n", calibrate(0, input))
	fmt.Printf("Part 2: %d\n", calibrateToFirstSeenTwice(0, input))
}

func calibrate(val int, adjustments []int) int {
	for _, adj := range adjustments {
		val += adj
	}

	return val
}

func calibrateToFirstSeenTwice(val int, adjustments []int) int {
	seen := map[int]bool{}

	for n := 0; !seen[val]; n = (n + 1) % len(adjustments) {
		seen[val] = true
		val += adjustments[n]
	}

	return val
}
