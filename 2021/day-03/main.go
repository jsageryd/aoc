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

	for _, bit := range bits(input) {
		ones, zeros := onesAndZeros(input, bit)
		if len(ones) > len(zeros) {
			rate += bit
		}
	}

	return rate
}

func epsilonRate(input []int) int {
	var rate int

	for _, bit := range bits(input) {
		ones, zeros := onesAndZeros(input, bit)
		if len(ones) < len(zeros) {
			rate += bit
		}
	}

	return rate
}

func oxygenGeneratorRating(input []int) int {
	keep := make([]int, len(input))
	copy(keep, input)

	for _, bit := range bits(input) {
		ones, zeros := onesAndZeros(keep, bit)

		if len(ones) < len(zeros) {
			keep = subtract(keep, ones)
		} else {
			keep = subtract(keep, zeros)
		}

		if len(keep) == 1 {
			break
		}
	}

	return keep[0]
}

func co2ScrubberRating(input []int) int {
	keep := make([]int, len(input))
	copy(keep, input)

	for _, bit := range bits(input) {
		ones, zeros := onesAndZeros(keep, bit)

		if len(ones) >= len(zeros) {
			keep = subtract(keep, ones)
		} else {
			keep = subtract(keep, zeros)
		}

		if len(keep) == 1 {
			break
		}
	}

	return keep[0]
}

func bits(input []int) []int {
	var bits []int

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

	for bit := max; bit >= 1; bit /= 2 {
		bits = append(bits, bit)
	}

	return bits
}

func onesAndZeros(input []int, bit int) (ones, zeros []int) {
	for _, n := range input {
		if n&bit != 0 {
			ones = append(ones, n)
		} else {
			zeros = append(zeros, n)
		}
	}

	return ones, zeros
}

func subtract(a, b []int) []int {
	for _, n := range b {
		for idx, m := range a {
			if n == m {
				a = append(a[:idx], a[idx+1:]...)
				break
			}
		}
	}
	return a
}
