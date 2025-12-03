package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input []string

	for s := bufio.NewScanner(os.Stdin); s.Scan(); {
		input = append(input, s.Text())
	}

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	var sum int

	for _, line := range input {
		bank := make([]int, len(line))

		for n := range line {
			bank[n] = int(line[n]) - '0'
		}

		joltA := bank[0]
		idxA := 0

		for idx := 1; idx < len(bank)-1; idx++ {
			if bank[idx] > joltA {
				joltA = bank[idx]
				idxA = idx
			}
		}

		joltB := bank[idxA+1]

		for idx := idxA + 1; idx < len(bank); idx++ {
			if bank[idx] > joltB {
				joltB = bank[idx]
			}
		}

		sum += joltA*10 + joltB
	}

	return sum
}

func part2(input []string) int {
	var sum int

	for _, line := range input {
		bank := make([]int, len(line))

		for n := range line {
			bank[n] = int(line[n]) - '0'
		}

		comb := maxComb(bank, 12)

		var total int

		m := 1

		for n := len(comb) - 1; n >= 0; n-- {
			total += comb[n] * m
			m *= 10
		}

		sum += total
	}

	return sum
}

// maxComb picks k elements from digits until it finds the combination with the
// highest numeric value formed by the picked digits.
func maxComb(digits []int, k int) []int {
	comb := make([]int, k)

	var rec func(dd []int, cc []int) []int

	rec = func(dd []int, cc []int) []int {
		var maxHead int

		for n := 0; n <= len(dd)-len(cc); n++ {
			if dd[n] <= maxHead {
				continue
			}

			maxHead = dd[n]

			cc[0] = dd[n]

			if len(cc) > 1 {
				rec(dd[n+1:], cc[1:])
			}
		}

		return cc
	}

	return rec(digits, comb)
}
