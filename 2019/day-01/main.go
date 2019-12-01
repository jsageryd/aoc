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
	fmt.Printf("Part 2: %d\n", totalFuelWithFuel(input))
}

func fuel(mass int) int {
	return mass/3 - 2
}

func fuelWithFuel(mass int) int {
	f := fuel(mass)

	for addedF := fuel(f); addedF > 0; addedF = fuel(addedF) {
		f += addedF
	}

	return f
}

func totalFuel(masses []int) int {
	var total int
	for _, mass := range masses {
		total += fuel(mass)
	}
	return total
}

func totalFuelWithFuel(masses []int) int {
	var total int
	for _, mass := range masses {
		total += fuelWithFuel(mass)
	}
	return total
}
