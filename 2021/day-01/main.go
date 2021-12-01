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

	fmt.Printf("Part 1: %d\n", countIncreases(input))
}

func countIncreases(input []int) int {
	var count int
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			count++
		}
	}
	return count
}
