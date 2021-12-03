package main

import (
	"fmt"
)

func main() {
	var input []int

	for {
		var n int
		if _, err := fmt.Scanf("%b", &n); err != nil {
			break
		}
		input = append(input, n)
	}

	fmt.Printf("Part 1: %d\n", gammaRate(input)*epsilonRate(input))
}

func gammaRate(input []int) int {
	var rate int

	var max int
	for _, n := range input {
		if n > max {
			max = n
		}
	}

	for bit := 1; bit < max; bit *= 2 {
		var ones, zeros int
		for _, n := range input {
			if n&bit != 0 {
				ones++
			} else {
				zeros++
			}
		}
		if ones > zeros {
			rate += bit
		}
	}

	return rate
}

func epsilonRate(input []int) int {
	var rate int

	var max int
	for _, n := range input {
		if n > max {
			max = n
		}
	}

	for bit := 1; bit < max; bit *= 2 {
		var ones, zeros int
		for _, n := range input {
			if n&bit != 0 {
				ones++
			} else {
				zeros++
			}
		}
		if ones < zeros {
			rate += bit
		}
	}

	return rate
}
