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

	fmt.Printf("Part 1: %d\n", totalFuel(input))
}

func fuel(mass int) int {
	return mass/3 - 2
}

func totalFuel(masses []int) int {
	var total int
	for _, mass := range masses {
		total += fuel(mass)
	}
	return total
}
