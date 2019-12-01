package main

import (
	"fmt"
)

func main() {
	var totalFuel, totalFuelWithFuel int

	for {
		var mass int
		if _, err := fmt.Scanln(&mass); err != nil {
			break
		}
		totalFuel += fuel(mass)
		totalFuelWithFuel += fuelWithFuel(mass)
	}

	fmt.Printf("Part 1: %d\n", totalFuel)
	fmt.Printf("Part 2: %d\n", totalFuelWithFuel)
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
