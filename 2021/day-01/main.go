package main

import "fmt"

func main() {
	var input []int

	for {
		var n int
		if _, err := fmt.Scanln(&n); err != nil {
			break
		}
		input = append(input, n)
	}

	fmt.Printf("Part 1: %d\n", countIncreases(input, 1))
	fmt.Printf("Part 2: %d\n", countIncreases(input, 3))
}

func countIncreases(input []int, windowSize int) int {
	var count int
	for i := windowSize; i < len(input); i++ {
		if input[i] > input[i-windowSize] {
			count++
		}
	}
	return count
}
