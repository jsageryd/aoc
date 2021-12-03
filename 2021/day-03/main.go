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
	fmt.Printf("Part 2: %d\n", oxygenGeneratorRating(input)*co2ScrubberRating(input))
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

func oxygenGeneratorRating(input []int) int {
	var max int
	for _, n := range input {
		if n > max {
			max = n
		}
	}

	for i := 1; ; i *= 2 {
		if i >= max {
			max = i / 2
			break
		}
	}

	keep := make(map[int]struct{})
	for _, n := range input {
		keep[n] = struct{}{}
	}

outer:
	for {
		for bit := max; bit >= 1; bit /= 2 {
			if len(keep) == 1 {
				break outer
			}

			var nOnes, nZeros []int

			for n := range keep {
				if n&bit != 0 {
					nOnes = append(nOnes, n)
				} else {
					nZeros = append(nZeros, n)
				}
			}

			var nRemove []int

			if len(nOnes) < len(nZeros) {
				nRemove = nOnes
			} else {
				nRemove = nZeros
			}

			for _, n := range nRemove {
				delete(keep, n)
			}
		}
	}

	var rating int

	for k := range keep {
		rating = k
	}

	return rating
}

func co2ScrubberRating(input []int) int {
	var max int
	for _, n := range input {
		if n > max {
			max = n
		}
	}

	for i := 1; ; i *= 2 {
		if i >= max {
			max = i / 2
			break
		}
	}

	keep := make(map[int]struct{})
	for _, n := range input {
		keep[n] = struct{}{}
	}

outer:
	for {
		for bit := max; bit >= 1; bit /= 2 {
			if len(keep) == 1 {
				break outer
			}

			var nOnes, nZeros []int

			for n := range keep {
				if n&bit != 0 {
					nOnes = append(nOnes, n)
				} else {
					nZeros = append(nZeros, n)
				}
			}

			var nRemove []int

			if len(nOnes) >= len(nZeros) {
				nRemove = nOnes
			} else {
				nRemove = nZeros
			}

			for _, n := range nRemove {
				delete(keep, n)
			}

		}
	}

	var rating int

	for k := range keep {
		rating = k
	}

	return rating
}
