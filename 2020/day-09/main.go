package main

import (
	"fmt"
	"sort"
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

	fmt.Printf("Part 1: %d\n", findNumber(input, 25))
	fmt.Printf("Part 2: %d\n", findWeakness(input, findNumber(input, 25)))
}

func findNumber(input []int, windowSize int) int {
next:
	for i := windowSize; i < len(input); i++ {
		for j := i - windowSize; j < i; j++ {
			for k := i - windowSize; k < i; k++ {
				if input[j] == input[k] {
					continue
				}
				if input[i] == input[j]+input[k] {
					continue next
				}
			}
		}
		return input[i]
	}
	return 0
}

func findWeakness(input []int, invalidNumber int) int {
	for i := 0; i < len(input); i++ {
		var sum int
		for j := i; j < len(input); j++ {
			sum += input[j]
			if sum == invalidNumber {
				sort.Ints(input[i : j+1])
				return input[i] + input[j]
			}
		}
	}
	return 0
}
