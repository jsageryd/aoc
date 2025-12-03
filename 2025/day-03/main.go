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
